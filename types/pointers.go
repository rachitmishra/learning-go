package types

import "fmt"

func Pointer() {
	i := 1

	fmt.Println("value at i: ", i)
	fmt.Println("pointer: ", &i)
	j := &i

	fmt.Println("j: ", j)
	fmt.Println("value at j: ", *j)
	fmt.Println("change value")
	*j = 2

	fmt.Println("value at i: ", i)
	fmt.Println("value at j: ", j)
}
