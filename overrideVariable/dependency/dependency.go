package dependency

import "fmt"

var HelloWorld = func() error {
	fmt.Println("Hello World")
	return nil
}
