package transport

import (
	"maps/route"
	"maps/user"
)

type Train struct {
	passengers []*user.Passenger
}

func (t *Train) TakePassenger(p *user.Passenger) {
	route.BasicGetPassenger(t, p)
}

func (t *Train) OutPassenger(p *user.Passenger) {
	route.BasicOutPassenger(t, p)
}

func (t *Train) GetPassengers() []*user.Passenger {
	return t.passengers
}

func (t *Train) SetPassengers(ps []*user.Passenger) {
	t.passengers = ps
}

func (t *Train) Run() {
	route.BasicRunTransport(t)
}
