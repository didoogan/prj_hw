package main

import "fmt"

type Animal struct {
	Breed    string
	Nickname string
}

type Zookeeper struct {
	Name string
}

func (z *Zookeeper) CatchAnimal(c *Cage, a *Animal) {
	c.Animals = append(c.Animals, a)
	fmt.Printf("%s: I catch %s into cage\n", z.Name, a.Nickname)
}

type Cage struct {
	Animals []*Animal
}

func (c Cage) GetCaptured() {
	animalsInsideCage := len(c.Animals)
	if animalsInsideCage == 0 {
		fmt.Println("No one in cage")
	} else {
		fmt.Printf("There %v animals inside a cage\n", animalsInsideCage)
	}
}

func main() {
	lion := Animal{Breed: "lion", Nickname: "Leo"}
	tiger := Animal{Breed: "tiger", Nickname: "Brandon"}
	tortle := Animal{Breed: "tortle", Nickname: "Matilda"}
	raccoon := Animal{Breed: "raccoon", Nickname: "Antony"}
	python := Animal{Breed: "snake", Nickname: "Guido"}

	cage := Cage{}
	jhon := Zookeeper{Name: "Dic"}

	cage.GetCaptured()

	var freeAnimals []*Animal
	freeAnimals = append(freeAnimals, &lion, &tiger, &tortle, &raccoon, &python)

	lion.Nickname = "Leonessa"
	for _, a := range freeAnimals {
		jhon.CatchAnimal(&cage, a)
	}

	cage.GetCaptured()
}
