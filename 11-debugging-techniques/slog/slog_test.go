package slog_test

import (
	"testing"

	"github.com/daggerok/SSWU-GO001/11-debugging-techniques/slog"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	t.Run("should log", func(t *testing.T) {
		res, err := slog.Divide(10, 0)
		assert.Equal(t, 0, res)
		assert.NotNil(t, err)

		res10, _ := slog.Divide(10, 1)
		assert.Equal(t, 10, res10)
		assert.NotNil(t, err)

		res5, _ := slog.Divide(10, 2)
		assert.Equal(t, 5, res5)
		assert.NotNil(t, err)
	})
}
