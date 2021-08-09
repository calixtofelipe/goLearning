package main

import (
	"errors"
	"fmt"
)

const (
	HOME_TEST = "TESTE_CONST"
	HOME_PROD = "PROD_CONST"
)

func main() {
	name, last, err := name("felipe", "calixtoshuazenager")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(name, last)
}

func name(name string, lastname string) (string, string, error) {
	if len(lastname) > 10 {
		return "", "", errors.New("Error lastname more 10 characters")
	}
	return name, lastname, nil

}
