package log_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/11-debugging-techniques/log"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	t.Run("should log", func(t *testing.T) {
		res, err := log.Divide(10, 0)
		assert.Equal(t, 0, res)
		assert.Nil(t, err)

		res10, _ := log.Divide(10, 1)
		assert.Equal(t, 10, res10)
		assert.Nil(t, err)

		res5, _ := log.Divide(10, 2)
		assert.Equal(t, 5, res5)
		assert.Nil(t, err)
	})
}
