package main

import "os/exec"

func main() {
	goimportsPath, err := exec.LookPath("goimports")
	print(goimportsPath)
	if err != nil {
		panic(err)
	}
}
