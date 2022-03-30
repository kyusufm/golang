package main

import (
	"apertama/calculation"
	"apertama/multiply"
	"fmt"
)

func main() {
	fmt.Println("In the name of God, May God help us with Go Lang")
	//sentence := testAja()
	//fmt.Println(sentence)

	result := multiply.Mult(8, 9)
	resultTambah := calculation.Add(8, 9)

	fmt.Println(result, resultTambah)

}
