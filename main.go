package main

import example "github.com/frenkoes/gowrapper-swig/number"

func main() {
	test := example.NewNumber()
	test.SetNumber(10)
	test.Print()
	example.DeleteNumber(test)
}
