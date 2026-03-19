package unit_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/10-testing-in-go/divide"
)

func TestUnitDivideByZero(t *testing.T) {
	res, err := divide.Divide(10, 0)
	if res != 0 {
		t.Error("Expected a nil, got ", res)
	}
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
