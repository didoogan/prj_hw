package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"
)

const roundDuration = 10

type Task struct {
	question    string
	answers     []string
	rightAnswer int
}

func (t Task) printQuestion() {
	fmt.Printf("\n%v\n", t.question)
	for i, a := range t.answers {
		fmt.Printf("%v) %v\n", i+1, a)
	}
	fmt.Println()
}

type User struct {
	name string
}

type UserAnswer struct {
	task           Task
	user           User
	answer         int
	answerDuration int
}

type Round struct {
	task        Task
	userAnswers map[User]int
	mx          sync.Mutex
}

func (r *Round) askQuestion(g *Game) {
	ctx, cancelFunc := context.WithCancel(g.ctx)
	defer cancelFunc()

	select {
	case <-ctx.Done():
		fmt.Println("askQuestion has been gracefully shut downed")
		cancelFunc()
	default:
		r.task.printQuestion()
		g.questionCh <- r.task
		time.Sleep(time.Second * roundDuration)
	}
}

func (r *Round) userAnswerQuestion(u User, t Task, g *Game) {
	g.wg.Add(1)
	defer g.wg.Done()

	answerDuration := rand.Intn(roundDuration)
	time.Sleep(time.Second * time.Duration(answerDuration))

	randAnswer := rand.Intn(len(r.task.answers))

	r.mx.Lock()
	r.userAnswers[u] = randAnswer
	r.mx.Unlock()

	userAnswer := UserAnswer{task: t, user: u, answer: randAnswer, answerDuration: answerDuration}
	g.answerCh <- userAnswer

	fmt.Printf("%v answered %v)%v (%v sec)\n", u.name, randAnswer+1, t.answers[randAnswer], answerDuration)

}

func (r *Round) usersAnswerQuestion(g *Game, ctx context.Context) {
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	g.wg.Add(1)
	defer g.wg.Done()

	defer close(g.questionCh)

	answerFunc := func(t Task, g *Game) {
		for _, u := range g.users {

			answerTimeout := rand.Intn(roundDuration)
			time.Sleep(time.Duration(answerTimeout))

			go r.userAnswerQuestion(u, t, g)
		}
	}
	for t := range g.questionCh {
		select {
		case <-ctx.Done():
			fmt.Println("Begining of greceful shut down of usersAnswerQuestion")
			answerFunc(t, g)
			fmt.Println("usersAnswerQuestion has been grecefully shut downed")
			cancelFunc()
		default:
			answerFunc(t, g)
		}
	}
}

type Game struct {
	name        string
	users       []User
	rounds      []*Round
	usersPoints map[User]float64
	mx          sync.Mutex
	wg          sync.WaitGroup
	ctx         context.Context
	questionCh  chan Task
	answerCh    chan UserAnswer
}

func (g *Game) updateResult() {
	ctx, cancelFunx := context.WithCancel(g.ctx)
	defer cancelFunx()

	select {
	case <-ctx.Done():
		fmt.Println("updateResult has been gracefully shut downed")
		return
	default:
		g.wg.Add(1)
		defer g.wg.Done()

		defer close(g.answerCh)

		for ua := range g.answerCh {
			u := ua.user

			g.mx.Lock()
			points, ok := g.usersPoints[ua.user]
			if ok {
				penalty := float64(ua.answerDuration) * 0.1
				g.usersPoints[u] = points + 1 - penalty
			} else {
				g.usersPoints[u] = 1.0
			}
			g.mx.Unlock()
		}
	}
}

func (g *Game) printResult() {
	ctx, cancelFunc := context.WithCancel(g.ctx)
	defer cancelFunc()

	select {
	case <-ctx.Done():
		fmt.Println("printResult has been gracefully shut downed")
		cancelFunc()
	default:
		fmt.Printf("\nResult of %v\n", g.name)

		var userPointsPairs []struct {
			user   User
			points float64
		}

		for user, points := range g.usersPoints {
			userPointsPairs = append(userPointsPairs, struct {
				user   User
				points float64
			}{user: user, points: points})
		}

		sort.Slice(userPointsPairs, func(i, j int) bool {
			return userPointsPairs[i].points > userPointsPairs[j].points
		})

		for i, upp := range userPointsPairs {
			fmt.Printf("%v. %v: %.2f\n", i+1, upp.user.name, upp.points)
		}
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
		stop()
	default:
		tasks := []Task{
			{
				question: "What is the primary purpose of the Go context package?",
				answers: []string{
					"To manage database connections",
					"To handle HTTP requests",
					"To carry deadlines, cancellation signals, and other request-scoped values",
					"To create goroutines",
				},
				rightAnswer: 2,
			},
			{
				question: "Which function creates a derived context with a deadline?",
				answers: []string{
					"WithTimeout",
					"WithCancel",
					"WithValue",
					"WithDeadline",
				},
				rightAnswer: 3,
			},
			{
				question: "What should you do if a function requires a Context parameter but you’re unsure which context to use?",
				answers: []string{
					"Pass a nil context",
					"Use context.TODO()",
					"Create a custom context",
					"Ignore the context",
				},
				rightAnswer: 1,
			},
		}

		userNames := []string{"Jhon", "Linda", "Mike", "Alice", "David", "Emily", "Henry", "Grace", "Oliver", "Sophia", "Daniel", "Isabella", "William", "Ava", "Samuel"}
		users := make([]User, len(userNames), len(userNames))

		for i, u := range userNames {
			users[i] = User{u}
		}

		game := Game{
			name:        "The worldwide kahoot chempionship",
			users:       users,
			rounds:      make([]*Round, len(tasks), len(tasks)),
			usersPoints: make(map[User]float64),
			questionCh:  make(chan Task),
			answerCh:    make(chan UserAnswer),
			wg:          sync.WaitGroup{},
			ctx:         ctx,
		}

		go game.updateResult()

		for i, t := range tasks {
			round := Round{task: t, userAnswers: make(map[User]int)}
			game.rounds[i] = &round

			go round.usersAnswerQuestion(&game, ctx)
			round.askQuestion(&game)
		}

		// game.wg.Wait()
		game.printResult()
	}

}
