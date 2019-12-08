package main

import (
	"fmt"
)

type Staff interface {
	work() Ingredient
}

type Person struct {
	name string
}

type Ingredient struct {
	name string
}

func (p Person) move() {
	fmt.Printf("%s moves\n", p.name)
}

type Chopper struct {
	Person
}

type RiceCooker struct {
	Person
}

type Wrapper struct {
	Person
}

func (chopper Chopper) work() Ingredient {
	chopper.move()
	fmt.Printf("%s chops cucumber\n", chopper.name)
	return Ingredient{name: "🥒 cucumber"}
}

func (riceCooker RiceCooker) work() Ingredient {
	riceCooker.move()
	fmt.Printf("%s cooks rice\n", riceCooker.name)
	return Ingredient{name: "🍚 rice"}
}

func (wrapper Wrapper) work() Ingredient {
	wrapper.move()
	fmt.Printf("%s wraps nori\n", wrapper.name)
	return Ingredient{name: "🍙 nori"}
}

func prepareSushi(staffs ...Staff) {
	for _, staff := range staffs {
		ingredient := staff.work()
		fmt.Printf("%s is ready\n\n", ingredient.name)
	}
	fmt.Println("🍣 Sushi is ready")
}

func main() {
	chopper := Chopper{Person{name: "Chris"}}
	riceCooker := RiceCooker{Person{name: "Rick"}}
	wrapper := Wrapper{Person{name: "Will"}}

	prepareSushi(chopper, riceCooker, wrapper)
}
