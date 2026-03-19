package divide_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/10-testing-in-go/divide"
	//"github.com/stretchr/testify/assert"
)

func BenchmarkDivide(b *testing.B) {
	for i := 1; i < b.N; i++ {
		divide.Divide(2 * i, 1 * i)
		//assert.NotNil(b, r)
		//assert.Nil(b, err)
	}
}
