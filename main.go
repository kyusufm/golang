package main

import (
	//"apertama/calculation"
	//"apertama/multiply"
	"fmt"
)

func main() {
	fmt.Println("In the name of God, May God help us with Go Lang")
	//sentence := testAja()
	//fmt.Println(sentence)

	//result := multiply.Mult(8, 9)
	//resultTambah := calculation.Add(8, 9)

	//fmt.Println(result, resultTambah)

	//var name string = "Golang"

	//var name string
	//name = "joki"
	//fmt.Println(name)

	// age := 21

	// if age > 20 {
	// 	fmt.Println("sudah dewasa?")
	// } else if age >= 10 {
	// 	fmt.Println("wah udah besar")
	// } else {
	// 	fmt.Println("wah masih kecil")
	// }

	//case
	numero := 2

	switch numero {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga")
	default:
		fmt.Println("tidak terpenuhi")
	}

	switch {
	case numero > 8:
		fmt.Println("kenceng nih")
	case numero > 4:
		fmt.Println("santai bos")
	case numero > 1:
		fmt.Println("pelan pelan")
	default:
		fmt.Println("jangan malas")
	}
}
