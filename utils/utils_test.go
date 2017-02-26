package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, Add(1, 2), 3, "they should be equal")
}
