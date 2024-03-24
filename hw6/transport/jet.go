package transport

import (
	"maps/user"
)

type Jet struct {
	passengers []*user.Passenger
}

func (j *Jet) GetPassenger(p *user.Passenger) {
	basicGetPassenger(j, p)
}

func (j *Jet) OutPassenger(p *user.Passenger) {
	basicOutPassenger(j, p)
}

func (j *Jet) GetPassengers() []*user.Passenger {
	return j.passengers
}

func (j *Jet) SetPassengers(ps []*user.Passenger) {
	j.passengers = ps
}

func (j *Jet) Run() {
	basicRunTransport(j)
}
