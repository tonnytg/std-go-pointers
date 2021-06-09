package main

import (
	"testing"
)

func TestNewPersonPositiveAge(t *testing.T) {
	name := "test"
	_, err := ShowName(&name)
	if err != nil {
		t.Errorf("Expected received %v", err)
	}
}