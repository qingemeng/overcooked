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

func prepareSushi(staffs ...Staff) {
	for _, staff := range staffs {
		ingredient := staff.work()
		fmt.Printf("%s is ready\n\n", ingredient.name)
	}
	fmt.Println("🍣 Sushi is ready")
}

func main() {
	//prepareSushi(chopper, riceCooker, wrapper)
}
