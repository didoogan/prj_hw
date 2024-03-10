package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type HealsAmount int

type User struct {
	Name      string
	Health    HealsAmount
	Inventory []*Object
}

func (u *User) AddObject(o *Object) {
	u.Inventory = append(u.Inventory, o)
}

func (u *User) HasUserObject(objectName string) bool {
	for _, obj := range u.Inventory {
		if obj.Name == objectName {
			return true
		}
	}
	return false
}

func (u *User) ChangeHealth(amount HealsAmount) {
	calculatedHealth := u.Health + amount

	if calculatedHealth < 100 {
		u.Health = calculatedHealth
	}
}

type Object struct {
	Name string
}

type Situation struct {
	Description string
	Quest       func(u *User) error
}

func quest1(u *User) error {
	fmt.Println("There are three object on the floor: knife (1), flashlight(2) and guidline for helicopters pilots beginers (3). Choose one of those: \n")

	var objectChose int
	var object Object
	fmt.Scanln(&objectChose)

	if objectChose > 3 {
		return errors.New("You entered wrong value")
	}

	switch objectChose {
	case 1:
		object.Name = "knife"
	case 2:
		object.Name = "flashlight"
	case 3:
		object.Name = "guide"
	}

	u.AddObject(&object)

	return nil
}

func runQuestWithRandomGues(u *User, objectName string, challenge string) error {
	var choiceRange int

	if u.HasUserObject(objectName) {
		fmt.Printf("You lucky you have %v\n", objectName)
		choiceRange = 5
	} else {
		fmt.Printf("It's sad you don't have %v\n", objectName)
		choiceRange = 20
	}

	randomNumber := rand.Intn(choiceRange + 1)

	guesNumber := 0

	var guesNumberStatus string

	for {
		fmt.Printf("The are a lot of %vs. Choose one from 1 to %v. But remember each atempt will take 10%% of your health:\n", challenge, choiceRange)
		fmt.Scanln(&guesNumber)
		if guesNumber == randomNumber {
			break
		} else if guesNumber > randomNumber {
			guesNumberStatus = "more"
		} else {
			guesNumberStatus = "less"
		}

		u.ChangeHealth(-10)

		if u.Health < 1 {
			return errors.New("Game over")
		}

		fmt.Printf("Guesed %v number is %v than required one. You have %v percents of health\n", challenge, guesNumberStatus, u.Health)

	}

	fmt.Printf("You pass this quest, congratilation, you have %v %% of health\n", u.Health)

	return nil
}

func quest2(u *User) error {
	error := runQuestWithRandomGues(u, "flashlight", "branch")
	if error != nil {
		return error
	}
	return nil
}

func quest3(u *User) error {
	error := runQuestWithRandomGues(u, "knife", "weak point")
	if error != nil {
		return error
	}
	return nil
}

func quest4(u *User) error {
	error := runQuestWithRandomGues(u, "guide", "tumbler")
	if error != nil {
		return error
	}
	return nil
}

func runSituations(ss []*Situation, u *User) error {
	if u.Health < 1 {
		return fmt.Errorf("Game over")
	}

	for _, s := range ss {
		fmt.Println(s.Description)
		error := s.Quest(u)
		if error != nil {
			return error
		}
	}

	return nil
}

func main() {
	fmt.Println("Hello user, input your name: ")

	var name string
	fmt.Scanln(&name)

	user := User{Name: name, Health: 100}

	situation1Desc := fmt.Sprintf("Hello, %v! You wake up on the uncharted island in dark cave. You should survive and get home. Good luck!\n", user.Name)
	situation1 := Situation{Description: situation1Desc, Quest: quest1}

	situation2Desc := "Ok, now you should get out from cave."
	situation2 := Situation{Description: situation2Desc, Quest: quest2}

	situation3Desc := "Ok, your path go through forest, but you met huge bear there. You should kill him!"
	situation3 := Situation{Description: situation3Desc, Quest: quest3}

	situation4Desc := "Ok, in the mountains you find out helicopter. You could be survived if you handle it!"
	situation4 := Situation{Description: situation4Desc, Quest: quest4}

	situations := []*Situation{&situation1, &situation2, &situation3, &situation4}

	error := runSituations(situations, &user)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%v, Congratilation! You win! You have %v %% of helath, don't forget to visit hospital\n", user.Name, user.Health)
	}
}
