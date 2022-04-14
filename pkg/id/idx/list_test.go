package idx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_Has(t *testing.T) {
	a := New[T]()
	b := New[T]()
	c := New[T]()
	l := List[T]{a, b}

	assert.True(t, l.Has(a))
	assert.True(t, l.Has(a, c))
	assert.False(t, l.Has(c))
	assert.False(t, List[T](nil).Has(a))
}

func TestList_Insert(t *testing.T) {
	a := New[T]()
	b := New[T]()
	c := New[T]()
	l := List[T]{a, b}

	assert.Equal(t, List[T]{a, b, c}, l.Insert(-1, c))
	assert.Equal(t, List[T]{c, a, b}, l.Insert(0, c))
	assert.Equal(t, List[T]{a, c, b}, l.Insert(1, c))
	assert.Equal(t, List[T]{a, b, c}, l.Insert(2, c))
	assert.Equal(t, List[T]{a, b, c}, l.Insert(3, c))
}
