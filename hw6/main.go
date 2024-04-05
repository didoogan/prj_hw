package main

import (
	"maps/route"
	"maps/transport"
	"maps/user"
)

func main() {
	nick := user.Passenger{Name: "Nick"}
	helen := user.Passenger{Name: "Helen"}

	bus := transport.Bus{}
	train := transport.Train{}
	jet := transport.Jet{}

	bus.TakePassenger(&nick)
	bus.TakePassenger(&helen)

	train.TakePassenger(&nick)
	train.TakePassenger(&helen)

	jet.TakePassenger(&nick)
	jet.TakePassenger(&helen)

	euroToure := route.EuroRoute{}

	euroToure.AddTransport(&bus)
	euroToure.AddTransport(&train)

	jet.OutPassenger(&nick)
	jet.OutPassenger(&nick)
	euroToure.AddTransport(&jet)

	euroToure.ShowTransports()
	euroToure.Run()
}
