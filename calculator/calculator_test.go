package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2,3) = %d, expected %d", result, expected)
	}
}

func TestSubstract(t *testing.T) {
	result := Subtract(5, 2)
	expected := 3

	if result != expected {
		t.Errorf("Substract(5, 2) = %d, expected %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 5)
	expected := 10

	if expected != result {
		t.Errorf("Multiply(2, 5) %d, expected %d", result, expected)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(20, 4)
	expected := 5

	if result != expected {
		t.Errorf("Divide(20, 4)%d, expected%d", result, expected)
	}
}
