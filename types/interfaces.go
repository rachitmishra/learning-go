package types

import "fmt"

// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
type Animal interface {
	Speak()
}

type Dog struct {
}

func (d Dog) Speak() {
	fmt.Println("Bow!")
}

type Cat struct {
}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func Interface() {
	animals := []Animal{Dog{}, Cat{}}
	for _, animal := range animals {
		animal.Speak()
	}

	Woof(Dog{})
}

func Woof(val interface{}) {
	// all types satisfy the empty interface.
	fmt.Print(val)
}
