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
	task           *Task
	user           *User
	answer         int
	answerDuration int
}

type Round struct {
	task        Task
	userAnswers map[*User]int
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
		g.questionCh <- &r.task
		time.Sleep(time.Second * roundDuration)
	}
}

func (r *Round) usersAnswerQuestion(g *Game) {
	t, ok := <-g.questionCh
	if ok {
		for _, u := range g.users {
			go func(u *User, t *Task, g *Game) {
				answerDuration := rand.Intn(roundDuration)
				time.Sleep(time.Second * time.Duration(answerDuration))

				randAnswer := rand.Intn(len(r.task.answers))

				r.mx.Lock()
				r.userAnswers[u] = randAnswer
				r.mx.Unlock()

				userAnswer := UserAnswer{task: t, user: u, answer: randAnswer, answerDuration: answerDuration}
				g.answerCh <- &userAnswer

				fmt.Printf("%v answered %v)%v (%v sec)\n", u.name, randAnswer+1, t.answers[randAnswer], answerDuration)
			}(u, t, g)
		}
	}
}

type Game struct {
	name        string
	users       []*User
	rounds      []*Round
	usersPoints map[*User]float64
	mx          sync.Mutex
	ctx         context.Context
	questionCh  chan *Task
	answerCh    chan *UserAnswer
}

func (g *Game) updateResult() {
	for ua := range g.answerCh {
		if ua.answer != ua.task.rightAnswer {
			continue
		}

		u := ua.user

		g.mx.Lock()
		points, ok := g.usersPoints[ua.user]
		penalty := float64(ua.answerDuration) * 0.1
		bonus := 1 - penalty

		if ok {
			g.usersPoints[u] = points + bonus
		} else {
			g.usersPoints[u] = bonus
		}
		g.mx.Unlock()
	}
}

func (g *Game) printResult() {
	fmt.Printf("\nResult of %v\n", g.name)

	var userPointsPairs []struct {
		user   *User
		points float64
	}

	for user, points := range g.usersPoints {
		userPointsPairs = append(userPointsPairs, struct {
			user   *User
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

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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
	users := make([]*User, len(userNames), len(userNames))

	for i, u := range userNames {
		users[i] = &User{u}
	}

	game := Game{
		name:        "The worldwide kahoot chempionship",
		users:       users,
		rounds:      make([]*Round, len(tasks), len(tasks)),
		usersPoints: make(map[*User]float64),
		questionCh:  make(chan *Task),
		answerCh:    make(chan *UserAnswer),
		ctx:         ctx,
	}

	defer close(game.questionCh)
	defer close(game.answerCh)

	go game.updateResult()

	for i, t := range tasks {
		round := Round{task: t, userAnswers: make(map[*User]int)}
		game.rounds[i] = &round

		go round.usersAnswerQuestion(&game)
		round.askQuestion(&game)
	}

	game.printResult()
}
