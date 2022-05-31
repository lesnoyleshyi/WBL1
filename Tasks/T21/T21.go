package main

import (
	"WBL1/Tasks/T21/duck"
	"WBL1/Tasks/T21/turkey"
	t "WBL1/Tasks/T21/turkeyAdapter"
	"fmt"
)

func main() {
	mallardDuck := &duck.MallardDuck{}
	wildTurkey := &turkey.WildTurkey{}
	ta := t.NewTurkeyAdapter(wildTurkey)

	fmt.Println("The Duck says...")
	testDuck(mallardDuck)

	fmt.Println("\nThe Turkey Adapter says...")
	testDuck(ta)
}

func testDuck(d duck.Duck) {
	d.Quack()
	d.Fly()
}