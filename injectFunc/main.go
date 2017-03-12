package main

import (
	"fmt"
)

func main() {
	fmt.Println("dummy")
}

func CallHelloWorld(helloWorld func() error) error {
	err := helloWorld()
	return err
}
