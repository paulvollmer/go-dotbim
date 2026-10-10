package dotbim_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/paulvollmer/go-dotbim"
)

func TestNewColor(t *testing.T) {
	color := dotbim.NewColor(255, 210, 0, 255)

	assert.Equal(t, uint8(255), color.R)
	assert.Equal(t, uint8(210), color.G)
	assert.Equal(t, uint8(0), color.B)
	assert.Equal(t, uint8(255), color.A)
}

func TestNewColorTransparent(t *testing.T) {
	color := dotbim.NewColor(0, 0, 0, 0)

	assert.Equal(t, &dotbim.Color{}, color)
}
