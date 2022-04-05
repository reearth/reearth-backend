package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	assert.True(t, All([]int{1, 2, 3}, func(i int) bool { return i < 4 }))
	assert.False(t, All([]int{1, 2, 3}, func(i int) bool { return i < 3 }))
}

func TestAny(t *testing.T) {
	assert.True(t, Any([]int{1, 2, 3}, func(i int) bool { return i == 1 }))
	assert.False(t, Any([]int{1, 2, 3}, func(i int) bool { return i == 4 }))
}
