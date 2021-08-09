package main

import (
	"fmt"
)

func main() {
	cap8slicedelete()

}

func cap8slice() {
	slice := []string{"banana", "pera", "uva"}
	for index, value := range slice {
		fmt.Printf("On index %v has value: %v.\n", index, value)
	}
}

func cap8slice2() {
	slice := []string{"banana", "pera", "uva"}
	twoElements := slice[0:2]
	fmt.Println(twoElements)
}

func cap8slicedelete() {
	//when we put ... we passing the slice items and not the slice.
	slice := []string{"banana", "pera", "uva"}
	slice = append(slice[0:2], slice[3:]...)

	fmt.Println(slice)
}
