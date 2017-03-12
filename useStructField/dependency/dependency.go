package dependency

import "fmt"

type DependencyInterface interface {
	HelloWorld() error
}

type DependencyStruct struct{}

func (d *DependencyStruct) HelloWorld() error {
	fmt.Println("Hello World")
	return nil
}
