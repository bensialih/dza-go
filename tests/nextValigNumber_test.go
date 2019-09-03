package tests

import (
	"dza-go/seeder"
	"testing"
)

func sum(number1 float64, number2 float64) float64 {
	return number1 + number2
}

func TestSum(t *testing.T) {
	result := seeder.NonZero(3.4, 0)
	if result != 3.4 {
		t.Errorf("NonZero was incorrect, got: %f, want: %f.", result, 3.4)
	}
}
