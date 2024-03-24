package transport

import (
	"maps/user"
)

type Bus struct {
	passengers []*user.Passenger
}

func (b *Bus) GetPassenger(p *user.Passenger) {
	basicGetPassenger(b, p)
}

func (b *Bus) OutPassenger(p *user.Passenger) {
	basicOutPassenger(b, p)
}

func (b *Bus) GetPassengers() []*user.Passenger {
	return b.passengers
}

func (b *Bus) SetPassengers(ps []*user.Passenger) {
	b.passengers = ps
}

func (b *Bus) Run() {
	basicRunTransport(b)
}
