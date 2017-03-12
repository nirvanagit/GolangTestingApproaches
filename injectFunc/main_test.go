package main

import (
	"fmt"
	"testing"
)

var expErr = fmt.Errorf("there was an error")

func myHelloWorld() error {
	return nil
}

func myHelloWorldErr() error {
	return expErr
}

func TestCallHelloWorld(t *testing.T) {

	err := CallHelloWorld(myHelloWorld)
	if err != nil {
		t.Errorf("expected error to be: nil, but got: %v", err)
	}
}

func TestCallHelloWorldErr(t *testing.T) {

	err := CallHelloWorld(myHelloWorldErr)
	if err.Error() != expErr.Error() {
		t.Errorf("expected error to be: %v, but got: %v", expErr, err)
	}

}
