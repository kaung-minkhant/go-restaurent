package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToFixed(t *testing.T) {
	var num float32 = 1.234
	output := ToFixed(num, 3)
	assert.Equal(t, float32(1.234), output)
}
