// Package beam builds a rectangular beam mesh.
package beam

import (
	"strconv"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

// Beam is a rectangular beam. Width is along X, length along Y, and depth
// along Z.
type Beam struct {
	Width  float64
	Length float64
	Depth  float64
}

// New returns a beam.
func New(width, length, depth float64) (Beam, error) {
	beam := Beam{Width: width, Length: length, Depth: depth}
	if _, err := beam.Mesh(); err != nil {
		return Beam{}, err
	}
	return beam, nil
}

// Type is the .bim element type for a beam.
func (b Beam) Type() string { return "Beam" }

// Info records the beam dimensions in meters.
func (b Beam) Info() map[string]string {
	return map[string]string{
		"Width":  format(b.Width),
		"Length": format(b.Length),
		"Depth":  format(b.Depth),
	}
}

// Mesh returns the beam as a box.
func (b Beam) Mesh() (*dotbim.Mesh, error) {
	return primitive.Box(b.Width, b.Length, b.Depth)
}

func format(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}
