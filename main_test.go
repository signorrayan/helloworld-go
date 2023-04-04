package main

import "testing"

func TestHelloWorld(t *testing.T) {
	if helloworld() != "Welcome to Devops Hobbies!" {
		t.Fatal("Test fail")
	}
}
