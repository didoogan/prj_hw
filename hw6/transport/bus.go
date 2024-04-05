package transport

import (
	"maps/route"
	"maps/user"
)

type Bus struct {
	passengers []*user.Passenger
}

func (b *Bus) TakePassenger(p *user.Passenger) {
	route.BasicGetPassenger(b, p)
}

func (b *Bus) OutPassenger(p *user.Passenger) {
	route.BasicOutPassenger(b, p)
}

func (b *Bus) GetPassengers() []*user.Passenger {
	return b.passengers
}

func (b *Bus) SetPassengers(ps []*user.Passenger) {
	b.passengers = ps
}

func (b *Bus) Run() {
	route.BasicRunTransport(b)
}
