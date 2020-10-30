package main

import (
	"fmt"

	"github.com/gitchander/pica/picalc/pigla"
)

func main() {
	testGLA()
	//testNOD()
}

func testGLA() {
	for i := 0; i < 6; i++ {
		pi := pigla.CalcPi(i)
		fmt.Printf("%d: %0.15f\n", i, pi)
	}
}

func testNOD() {
	re := pigla.NumberOfDigits(10, 127)
	fmt.Println(re)
}
