package main

import "fmt"

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func main() {
	profile := MakeProfile("Goku")
	fmt.Println("Name :", profile.Name)
	fmt.Println("Healt :", profile.Health)
	fmt.Println("Power :", profile.Power)
	fmt.Println("Exp :", profile.Exp)
	fmt.Println("=== HEROES POWER UP ===")

	powerUp := PowerUp(profile, 2)
	fmt.Println("Name :", powerUp.Name)
	fmt.Println("Healt :", powerUp.Health)
	fmt.Println("Power :", powerUp.Power)
	fmt.Println("Exp :", powerUp.Exp)

}

func MakeProfile(Name string) Profile {
	var profile Profile

	if Name == "sasuke" {
		profile.Name = "Sasuke"
		profile.Health = 200
		profile.Power = 100
		profile.Exp = 0
	} else if Name == "Naruto" {
		profile.Name = "Naruto"
		profile.Health = 150
		profile.Power = 200
		profile.Exp = 50
	} else if Name == "Goku" {
		profile.Name = "Goku"
		profile.Health = 400
		profile.Power = 300
		profile.Exp = 100
	}

	return profile
}

func PowerUp(p Profile, lalala int) Profile {
	p.Health = p.Health + (p.Health * lalala)
	p.Power = p.Power + (p.Power * lalala)
	p.Exp = p.Exp + (p.Exp * lalala)
	return p
}
