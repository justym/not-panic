package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("I'll be Panic !")

	err := nil
	if err != nil {
		panic(err)
	}

	err = nil 
	if err != nil {
		panic(err)
	}

	err = errors.New("Use Panic")
	bePanic(err)
}

func bePanic(err error){
	if err != nil {
		panic(err)
	}
}

