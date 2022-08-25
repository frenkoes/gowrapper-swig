package main

import "github.com/frenkoes/gowrapper-swig/number"

func main() {
	test := number.NewNumber()
	test.SetNumber(10)
	test.Print()
	number.DeleteNumber(test)
}
