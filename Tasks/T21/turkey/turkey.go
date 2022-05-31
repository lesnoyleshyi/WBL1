package turkey

import "fmt"

type Turkey interface {
	Gobble()
	Fly()
}
type WildTurkey struct {}

func (w *WildTurkey) Gobble()  {
	fmt.Println("Gobble Gobble")
}

func (w *WildTurkey) Fly()  {
	fmt.Println("I'm flying a short distance")
}
