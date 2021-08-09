package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Vehicle Vehicle
}

type Vehicle interface {
	buzz()
}

type Car struct {
	Branch  string
	Factory string
}

type Motorcycle struct {
	Branch         string
	Factory        string
	Displancements int
}

func (c Car) buzz() {
	fmt.Println(c.Factory, "buzzzzz!!")
}

func (m Motorcycle) buzz() {
	fmt.Println(m.Factory, "buzzzzz!!")
}

//if we use pointer(*) when we alter the p.Nome, we will alter for global scope.
//because we altering the value on memory address.
func (p *Person) run() {
	p.Name = "Felipe pointer"
	motorcycle := Motorcycle{
		Branch:         "bmw motorcycle 3",
		Factory:        "bmw11",
		Displancements: 10,
	}
	p.Vehicle = motorcycle
	fmt.Println(p.Name, "running")
	p.Vehicle.buzz()
}

func main() {
	bmw := Car{
		Branch:  "bmw v3",
		Factory: "bmw",
	}

	felipe := Person{
		Name:    "Felipe Calixto",
		Age:     35,
		Vehicle: bmw,
	}

	felipe.run()
	fmt.Println(felipe.Name, felipe.Age)
	felipe.Vehicle.buzz()
	//fmt.Println("Nome:  ", felipe.Name, felipe.Car.Branch)
}
