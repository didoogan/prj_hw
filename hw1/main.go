package main

import (
	"fmt"
	"time"
)

func main() {

	firstName := "Stepan"
	patronymicName := "Andriyovych"
	lastName := "Bandera"
	fullName := fmt.Sprintf("%s %s %s", firstName, patronymicName, lastName)
	dob := time.Date(1909, 1, int(time.January), 0, 0, 0, 0, time.Local)
	dobStr := fmt.Sprintf("%d %s %d", dob.Day(), dob.Month(), dob.Year())

	content := fullName + " was born on " + dobStr
	fmt.Println(content)
}
