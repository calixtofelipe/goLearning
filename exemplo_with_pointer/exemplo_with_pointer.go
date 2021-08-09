package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Car  Car
}

type Car struct {
	Branch  string
	Factory string
}

//if we use pointer(*) when we alter the p.Nome, we will alter for global scope.
//because we altering the value on memory address.
func (p *Person) run() {
	p.Name = "Felipe pointer"
	fmt.Println(p.Name, "running", p.Car.Factory)
}

func main() {
	felipe := Person{
		Name: "Felipe Calixto",
		Age:  35,
		Car: Car{
			Branch:  "i30",
			Factory: "Hyundai",
		},
	}

	felipe.run()
	fmt.Println(felipe.Name)
	//fmt.Println("Nome:  ", felipe.Name, felipe.Car.Branch)
}
