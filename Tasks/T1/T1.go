package main

import "fmt"

type Human struct {
	Name   string
	height uint8
	weight float32
}

func (h *Human) eat(foodWeightG float32) {
	h.weight += foodWeightG / 1000
}

func (h Human) Greet() {
	fmt.Printf("Hi! My name is %s\n", h.Name)
}

type Action struct {
	Human
	energyCost int32
}

type Action2 struct {
	h          Human
	energyCost int32
}

func main() {
	run := Action{
		Human{
			Name:   "Oleg",
			height: 150,
			weight: 75.5,
		},
		-100.0,
	}
	sleep := Action2{h: Human{"Oleg", 150, 75.5}, energyCost: 300}

	run.eat(1000)
	sleep.h.Greet()
	fmt.Println(run.weight)
}
