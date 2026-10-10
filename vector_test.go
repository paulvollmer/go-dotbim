package dotbim_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/paulvollmer/go-dotbim"
)

func TestNewVector(t *testing.T) {
	vector := dotbim.NewVector(1.5, -2, 3)

	assert.Equal(t, &dotbim.Vector{X: 1.5, Y: -2, Z: 3}, vector)
}
