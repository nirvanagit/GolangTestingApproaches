package main

import (
	"fmt"

	"github.com/nirvanagit/LearningGo/Blog/Testing/overrideVariable/dependency"
)

func main() {
	fmt.Println("dummy")
}

func CallHelloWorld() error {
	err := dependency.HelloWorld()
	return err
}
