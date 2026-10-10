package dotbim_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/paulvollmer/go-dotbim"
)

func TestNewRotation(t *testing.T) {
	rotation := dotbim.NewRotation(0.1, 0.2, 0.3, 0.4)

	assert.Equal(t, &dotbim.Rotation{X: 0.1, Y: 0.2, Z: 0.3, W: 0.4}, rotation)
}

func TestNewRotationIdentity(t *testing.T) {
	rotation := dotbim.NewRotation(0, 0, 0, 1)

	assert.Equal(t, &dotbim.Rotation{W: 1}, rotation)
}
