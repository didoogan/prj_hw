package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Student struct {
	name         string
	averageGrade int
}

func (s *Student) getInfo() string {
	return fmt.Sprintf("Student %v has average grade %v\n", s.name, s.averageGrade)
}

type Class struct {
	name     string
	students []*Student
}

func (c *Class) getInfo() string {
	return fmt.Sprintf("Class %v has %v students\n", c.name, len(c.students))
}

var studentNames = []string{"Patricia Halford", "Patrick Keating", "Bernard Szczepanek", "Joe Williams", "Andy Smith", "Christina Johnson", "John Johnson", "Andy Williams", "Joe Johnson", "John Smith"}
var students = make([]*Student, len(studentNames), len(studentNames))

const maxGrade = 12

var myClass Class

func main() {

	for i, sn := range studentNames {
		randomGrade := rand.Intn(maxGrade) + 1
		students[i] = &Student{name: sn, averageGrade: randomGrade}
	}

	myClass = Class{name: "11B", students: students}

	mux := http.NewServeMux()

	mux.HandleFunc("/", authorize(classInfoHandler))
	mux.HandleFunc("/student/{studentId}", authorize(studentInfoHandler))
	mux.HandleFunc("/login", loginPageHandler)

	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrAbortHandler {
		log.Printf("server serve failed: %s", err)
	}
}

func classInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, myClass.getInfo())
}

func studentInfoHandler(w http.ResponseWriter, r *http.Request) {
	i, err := strconv.Atoi(r.PathValue("studentId"))
	if err != nil || i < 0 || i >= len(studentNames) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, students[i].getInfo())
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		html := `<body><h4>Login Page</h4>
				<p>
				<form method="POST">
					<label for="login">Login:</label>
					<input name="login" type="text"><br>
					<label for="password">Password:</label>
					<input name="password" type="password"><br>
					<input type="submit" value="Submit">
				</form>
				</p>
			<body>`
		fmt.Fprint(w, html)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	token, ok := authenticateUser(login, password)
	if ok {
		http.SetCookie(w, &http.Cookie{
			Name:    sessionTokenCookieName,
			Value:   token,
			Expires: time.Now().Add(time.Minute),
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// not authenticated, show login form again
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
