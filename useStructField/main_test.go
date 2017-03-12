package main

import (
	"fmt"
	"testing"
)

var expErr = fmt.Errorf("there was an error")

type MockDependencyStruct struct {
	Err error
}

func (m *MockDependencyStruct) HelloWorld() error {
	return m.Err
}

func TestCallHelloWorld(t *testing.T) {

	initializeStruct := &myStruct{
		dependencies: &MockDependencyStruct{
			Err: nil,
		},
	}
	err := initializeStruct.CallHelloWorld()
	if err != nil {
		t.Errorf("expected error to be: nil, but got: %v", err)
	}
}

func TestCallHelloWorldErr(t *testing.T) {

	initializeStruct := &myStruct{
		dependencies: &MockDependencyStruct{
			Err: expErr,
		},
	}
	err := initializeStruct.CallHelloWorld()
	if err.Error() != expErr.Error() {
		t.Errorf("expected error to be: %v, but got: %v", expErr, err)
	}
}
