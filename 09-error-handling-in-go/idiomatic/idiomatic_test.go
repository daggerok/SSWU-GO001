package idiomatic_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/09-error-handling-in-go/idiomatic"
	"github.com/stretchr/testify/assert"
)

func TestIdiomaticGoWay(t *testing.T) {
	t.Run("should not accept negative values", func(t *testing.T) {
		_, err := idiomatic.DoSomething(-1)
		t.Log(err)
		assert.NotNil(t, err)
		assert.Equal(t, "input cannot be negative", err.Error())
	})

	t.Run("should not accept zero values", func(t *testing.T) {
		_, err := idiomatic.DoSomething(0)
		t.Log(err)
		assert.NotNil(t, err)
		assert.Equal(t, "input cannot be zero", err.Error())
	})

	t.Run("should accept positive values", func(t *testing.T) {
		res, err := idiomatic.DoSomething(1)
		t.Log(res)
		assert.Nil(t, err)
		assert.Equal(t, "Valid input: 1", res)
	})
}
