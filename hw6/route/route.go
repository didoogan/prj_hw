package route

import (
	"fmt"
)

type Route interface {
	AddTransport(t Transport)
	ShowTransports()
	GetTransports() []Transport
	Run()
}

func basicShowTransport(r Route) {
	fmt.Printf("During %T you will use the next transport:\n", r)
	for _, t := range r.GetTransports() {
		fmt.Printf("%T\n", t)
	}
}

func basicRun(r Route) {
	for _, t := range r.GetTransports() {
		t.Run()
	}
}
