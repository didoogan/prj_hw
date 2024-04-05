package transport

import (
	"maps/route"
	"maps/user"
)

type Jet struct {
	passengers []*user.Passenger
}

func (j *Jet) TakePassenger(p *user.Passenger) {
	route.BasicGetPassenger(j, p)
}

func (j *Jet) OutPassenger(p *user.Passenger) {
	route.BasicOutPassenger(j, p)
}

func (j *Jet) GetPassengers() []*user.Passenger {
	return j.passengers
}

func (j *Jet) SetPassengers(ps []*user.Passenger) {
	j.passengers = ps
}

func (j *Jet) Run() {
	route.BasicRunTransport(j)
}
