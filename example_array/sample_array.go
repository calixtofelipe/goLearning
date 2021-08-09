package main

import (
	"errors"
	"fmt"
)

func main() {
	//array fixed positions
	x := [10]int{4, 6, 7, 8, 10}
	fmt.Println(x)

	slice := make([]int, 5)
	slice[0] = 10
	slice[1] = 10
	slice[2] = 10
	slice[3] = 10
	slice[4] = 10
	fmt.Println(slice)
	slice = append(slice, 5)
	fmt.Println(slice)

	sliceString := []string{
		"string1", "string2",
	}
	fmt.Println(sliceString)

	sliceInt := []int{}

	fmt.Println(append(sliceInt, 1))

	mapsample := map[string]int{}
	mapsample["felipe"] = 35
	mapsample["ana"] = 29
	fmt.Println(mapsample["felipe"])

}

func name(name string, lastname string) (string, string, error) {
	if len(lastname) > 10 {
		return "", "", errors.New("Error lastname more 10 characters")
	}
	return name, lastname, nil

}
