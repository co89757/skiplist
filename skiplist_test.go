package skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddKey(t *testing.T) {
	assert := assert.New(t)
	list := New()
	list.Add("a", 1)
	v, _ := list.Get("a")
	assert.Equal(v, 1, "a should map to 1")
}

func TestRemoveKey(t *testing.T) {
	list := New()
	list.Add("a", 1)
	list.Add("b", 2)
	ok := list.Remove("a")
	assert.True(t, ok, "key should exist")
	_, exist := list.Get("a")
	assert.False(t, exist, "key should be removed")
}
