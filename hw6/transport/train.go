package transport

import (
	"maps/user"
)

type Train struct {
	passengers []*user.Passenger
}

func (t *Train) GetPassenger(p *user.Passenger) {
	basicGetPassenger(t, p)
}

func (t *Train) OutPassenger(p *user.Passenger) {
	basicOutPassenger(t, p)
}

func (t *Train) GetPassengers() []*user.Passenger {
	return t.passengers
}

func (t *Train) SetPassengers(ps []*user.Passenger) {
	t.passengers = ps
}

func (t *Train) Run() {
	basicRunTransport(t)
}
