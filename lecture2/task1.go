package main

import "fmt"

type Car struct {
	Brand string
	Color string
}

type Parking struct {
	Car *Car
}

func (p Parking) IsPlaceIsFree() bool {
	return p.Car == nil
}

func main() {

	toyota := Car{Brand: "toyota", Color: "white"}

	parking := Parking{Car: &toyota}

	var status string

	if parking.IsPlaceIsFree() {
		status = "available"
	} else {
		status = fmt.Sprintf("ocupied by %v %v", parking.Car.Color, parking.Car.Brand)
	}

	fmt.Printf("The place is %v\n", status)
}
