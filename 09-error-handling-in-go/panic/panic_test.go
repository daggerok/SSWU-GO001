package panic_test

import (
	"testing"

	ohMyPanic "github.com/daggerok/SSWU-GO001/09-error-handling-in-go/panic"
	"github.com/stretchr/testify/assert"
)

func TestPanicAndRecover(t *testing.T) {
	t.Run("should panic from divide by zero", func(t *testing.T) {
		t.Log("Expecting panic from divide by zero")

		assert.Panics(t, func() {
			ohMyPanic.DivideAndPanic(10, 0)
		}, "divide by zero")
	})

	t.Run("should panic and recover from divide by zero", func(t *testing.T) {
		t.Log("Expecting panic and recover from divide by zero")

		res := ohMyPanic.DividePanicAndRecover(10, 0)
		t.Log(res)
		assert.NotNil(t, res)
		assert.Equal(t, 0, res)
	})
}
