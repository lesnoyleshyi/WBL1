package turkeyAdapter

import "WBL1/Tasks/T21/turkey"

type turkeyAdapter struct {
	turkey turkey.Turkey
}

func NewTurkeyAdapter(t turkey.Turkey) *turkeyAdapter {
	return &turkeyAdapter{
		turkey: t,
	}
}

func (t *turkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t *turkeyAdapter) Fly() {
	for i := 1; i <= 5; i++ {
		t.turkey.Fly()
	}
}
