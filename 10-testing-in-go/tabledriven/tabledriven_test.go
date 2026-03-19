package tabledriven_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/10-testing-in-go/divide"
)

func TestTableDivide(t *testing.T) {
	tests := []struct {
		a, b int
		expected int
		wantErr bool
	}{
		{a: 10, b: 0, expected: 0, wantErr: true},
		{a: 10, b: 2, expected: 5, wantErr: false},
		{a: 10, b: 10, expected: 1, wantErr: false},
	}
	for _, test := range tests {
		t.Log(test)
		result, err := divide.Divide(test.a, test.b)
		t.Log(result, err)

		if (err != nil) != test.wantErr {
			t.Errorf("Divide(%d, %d): unexpected error status; got %v, wantErr %v", test.a, test.b, err != nil, test.wantErr)
		}
		if result != test.expected {
			t.Errorf("Divide(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}
