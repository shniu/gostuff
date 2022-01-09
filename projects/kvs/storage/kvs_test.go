package storage

import (
	"fmt"
	"testing"
)

type Flyable interface {
	Fly()
	Hello()
}

type Duck struct {
	flyable bool
}

func (d Duck) Hello() {
	fmt.Println("When i am flying, i will say hello")
}

func (d *Duck) Fly() {
	d.flyable = false
	fmt.Println("I am a duck, and can't fly")
}

func fly(f *Flyable) {
	(*f).Fly()
}

func TestOpen(t *testing.T) {
	var f Flyable = &Duck{flyable: true}
	f.Fly()
	f.Hello()
}
