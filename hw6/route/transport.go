package route

import (
	"fmt"
	"maps/user"
)

type Transport interface {
	TakePassenger(p *user.Passenger)
	OutPassenger(p *user.Passenger)

	GetPassengers() []*user.Passenger
	SetPassengers([]*user.Passenger)

	Run()
}

func BasicOutPassenger(t Transport, p *user.Passenger) {
	var removedPassengerIndex *int

	for index, passenger := range t.GetPassengers() {
		if passenger.Name == p.Name {
			removedPassengerIndex = &index
			break
		}
	}

	if removedPassengerIndex != nil {
		t.SetPassengers(append(t.GetPassengers()[:*removedPassengerIndex], t.GetPassengers()[*removedPassengerIndex+1:]...))
	} else {
		fmt.Printf("%T doesn't run passenger %v\n", t, p.Name)
	}
}

func BasicGetPassenger(t Transport, p *user.Passenger) {
	t.SetPassengers(append(t.GetPassengers(), p))
}

func BasicRunTransport(t Transport) {
	fmt.Printf("%T run passengers: \n", t)

	for _, p := range t.GetPassengers() {
		fmt.Println(p.Name)
	}
}
