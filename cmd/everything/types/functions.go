package types

import "fmt"

// https://gobyexample.com/functions
func Function() {
	fmt.Println("A Function")
	a, err := FunctionWithMultipleReturn()
	fmt.Println(a, err)
}

// https://gobyexample.com/multiple-return-values
func FunctionWithMultipleReturn() (int, error) {
	fmt.Println("Another function")
	return 0, nil
}

// https://gobyexample.com/variadic-functions
func FunctionWithVariableArgument(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
