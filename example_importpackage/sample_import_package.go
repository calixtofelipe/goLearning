package main

import (
	"fmt"
	"go-learning/packagesample"
)

const (
	HOME_TEST = "TESTE_CONST"
	HOME_PROD = "PROD_CONST"
)

func main() {
	packagesample.PrintExemplo()

	a := 10
	b := "test"
	c := 10.22
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)

	x := 10
	//y pointer to x
	y := &x
	//GET MEMORY ADDRESS
	fmt.Println("GET x MEMORY ADDRESS: ", &x)
	fmt.Println("GET y MEMORY ADDRESS: ", y)
	//GET MEMORY VALUE
	fmt.Println("GET the value on Y MEMORY ADDRESS: ", *y)

	//create z pointer
	var z *int = &x
	fmt.Println("GET z memory: ", z)
	fmt.Println("GET z value: ", *z)
}
