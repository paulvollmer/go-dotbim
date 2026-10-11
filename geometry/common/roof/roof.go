// Package roof builds a single-slope roof mesh.
package roof

import (
	"strconv"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

// Roof is a single slope. Width is along X and depth along Y. The rise is the
// full height at y = 0 and falls to zero at y = depth.
type Roof struct {
	Width float64
	Depth float64
	Rise  float64
}

// New returns a roof.
func New(width, depth, rise float64) (Roof, error) {
	roof := Roof{Width: width, Depth: depth, Rise: rise}
	if _, err := roof.Mesh(); err != nil {
		return Roof{}, err
	}
	return roof, nil
}

// Type is the .bim element type for a roof.
func (r Roof) Type() string { return "Roof" }

// Info records the roof dimensions in meters.
func (r Roof) Info() map[string]string {
	return map[string]string{
		"Width": format(r.Width),
		"Depth": format(r.Depth),
		"Rise":  format(r.Rise),
	}
}

// Mesh returns the roof as a wedge.
func (r Roof) Mesh() (*dotbim.Mesh, error) {
	return primitive.Wedge(r.Width, r.Depth, r.Rise)
}

func format(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}
