package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	a := &struct{}{}
	err := errors.New("ERR")
	assert.Same(t, a, Must(a, nil))
	assert.PanicsWithValue(t, err, func() {
		_ = Must(a, err)
	})
}
