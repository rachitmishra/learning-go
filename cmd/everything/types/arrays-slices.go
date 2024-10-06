package types

import "fmt"

func Arrays() {
	var a [5]int
	fmt.Println("emp: ", a)
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))
	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println("b: ", b)
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println("c: ", c)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}

func Slices() {

}
