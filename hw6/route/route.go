package route

import (
	"fmt"
	"maps/transport"
	"maps/utils"
)

type Route interface {
	AddTransport(t transport.Transport)
	ShowTransports()
	GetTransports() []transport.Transport
	Run()
}

func basicShowTransport(r Route) {
	fmt.Printf("During %v you will use the next transport:\n", utils.GetTypeName(r))
	for _, t := range r.GetTransports() {
		fmt.Println(utils.GetTypeName(t))
	}
}

func basicRun(r Route) {
	for _, t := range r.GetTransports() {
		t.Run()
	}
}
