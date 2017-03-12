package main

import (
	"fmt"

	"github.com/nirvanagit/LearningGo/Blog/Testing/useStructField/dependency"
)

func main() {
	fmt.Println("dummy")
}

type myStruct struct {
	dependencies dependency.DependencyInterface
}

func (m *myStruct) CallHelloWorld() error {
	err := m.dependencies.HelloWorld()
	return err
}
