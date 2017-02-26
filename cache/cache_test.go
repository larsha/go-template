package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	err := Set("foo", "bar")
	assert.Nil(t, err)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)

	val, err := Get("foo")
	assert.Nil(err)
	assert.Equal(val, "bar", "they should be equal")
}

func TestDel(t *testing.T) {
	err := Del("foo")
	assert.Nil(t, err)
}
