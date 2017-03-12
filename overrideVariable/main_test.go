package main

import (
	"fmt"
	"testing"

	"github.com/nirvanagit/LearningGo/Blog/Testing/overrideVariable/dependency"
)

func TestCallHelloWorld(t *testing.T) {
	var originalHelloWorld = dependency.HelloWorld
	defer func() { dependency.HelloWorld = originalHelloWorld }()

	dependency.HelloWorld = func() error {
		return nil
	}

	err := CallHelloWorld()
	if err != nil {
		t.Errorf("expected error to be: nil, but got: %v", err)
	}
}

func TestCallHelloWorldErr(t *testing.T) {
	var originalHelloWorld = dependency.HelloWorld
	var expErr = fmt.Errorf("there was an error")
	defer func() { dependency.HelloWorld = originalHelloWorld }()

	dependency.HelloWorld = func() error {
		return expErr
	}

	err := CallHelloWorld()
	if err.Error() != expErr.Error() {
		t.Errorf("expected error to be: %v, but got: %v", expErr, err)
	}
}
