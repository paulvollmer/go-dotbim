// Package slab builds a horizontal slab mesh.
package slab

import (
	"strconv"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

// Slab is a horizontal plate. Width is along X, depth along Y, and thickness
// along Z.
type Slab struct {
	Width     float64
	Depth     float64
	Thickness float64
}

// New returns a slab.
func New(width, depth, thickness float64) (Slab, error) {
	slab := Slab{Width: width, Depth: depth, Thickness: thickness}
	if _, err := slab.Mesh(); err != nil {
		return Slab{}, err
	}
	return slab, nil
}

// Type is the .bim element type for a slab.
func (s Slab) Type() string { return "Slab" }

// Info records the slab dimensions in meters.
func (s Slab) Info() map[string]string {
	return map[string]string{
		"Width":     format(s.Width),
		"Depth":     format(s.Depth),
		"Thickness": format(s.Thickness),
	}
}

// Mesh returns the slab as a box.
func (s Slab) Mesh() (*dotbim.Mesh, error) {
	return primitive.Box(s.Width, s.Depth, s.Thickness)
}

func format(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}
