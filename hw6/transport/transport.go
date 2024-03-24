package transport

import (
	"fmt"
	"maps/user"
	"maps/utils"
)

type Transport interface {
	GetPassenger(p *user.Passenger)
	OutPassenger(p *user.Passenger)

	GetPassengers() []*user.Passenger
	SetPassengers([]*user.Passenger)

	Run()
}

func basicOutPassenger(t Transport, p *user.Passenger) {
	var removedPassengerIndex int

	for index, passenger := range t.GetPassengers() {
		if passenger.Name == p.Name {
			removedPassengerIndex = index
			break
		}
	}
	t.SetPassengers(append(t.GetPassengers()[:removedPassengerIndex], t.GetPassengers()[removedPassengerIndex+1:]...))
}

func basicGetPassenger(t Transport, p *user.Passenger) {
	t.SetPassengers(append(t.GetPassengers(), p))
}

func basicRunTransport(t Transport) {
	fmt.Printf("%v run passengers: \n", utils.GetTypeName(t))

	for _, p := range t.GetPassengers() {
		fmt.Println(p.Name)
	}
}
