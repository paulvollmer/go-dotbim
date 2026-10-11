// Package door builds a door leaf mesh.
package door

import (
	"strconv"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

// Door is a rectangular door leaf. Width runs along Y, thickness along X, and
// height along Z.
type Door struct {
	Width     float64
	Height    float64
	Thickness float64
}

// New returns a door leaf.
func New(width, height, thickness float64) (Door, error) {
	door := Door{Width: width, Height: height, Thickness: thickness}
	if _, err := door.Mesh(); err != nil {
		return Door{}, err
	}
	return door, nil
}

// Type is the .bim element type for a door.
func (d Door) Type() string { return "Door" }

// Info records the door dimensions in meters.
func (d Door) Info() map[string]string {
	return map[string]string{
		"Width":     format(d.Width),
		"Height":    format(d.Height),
		"Thickness": format(d.Thickness),
	}
}

// Mesh returns the door leaf as a box.
func (d Door) Mesh() (*dotbim.Mesh, error) {
	return primitive.Box(d.Thickness, d.Width, d.Height)
}

func format(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}
