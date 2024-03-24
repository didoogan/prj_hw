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

	bus.GetPassenger(&nick)
	bus.GetPassenger(&helen)

	train.GetPassenger(&nick)
	train.GetPassenger(&helen)

	jet.GetPassenger(&nick)
	jet.GetPassenger(&helen)

	euroToure := route.EuroRoute{}

	euroToure.AddTransport(&bus)
	euroToure.AddTransport(&train)

	jet.OutPassenger(&nick)
	euroToure.AddTransport(&jet)

	euroToure.ShowTransports()
	euroToure.Run()
}
