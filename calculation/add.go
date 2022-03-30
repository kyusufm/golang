package calculation

//diawali huruf kapital agar bisa diakses dari package lain, karena beda nama package.
func Add(number int, numberTwo int) int {
	return add(number, numberTwo)
}

func add(number int, numberTwo int) int {
	return number + 1 + numberTwo
}
