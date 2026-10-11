// Package column builds rectangular and circular column meshes.
package column

import (
	"strconv"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

// Column is a vertical column. Height is along Z.
type Column struct {
	Width    float64
	Depth    float64
	Height   float64
	Radius   float64
	Segments int
	Circular bool
}

// New returns a rectangular column. Width is along X and depth along Y.
func New(width, depth, height float64) (Column, error) {
	column := Column{Width: width, Depth: depth, Height: height}
	if _, err := column.Mesh(); err != nil {
		return Column{}, err
	}
	return column, nil
}

// NewCircular returns a circular column. The base is centered on (radius, radius).
func NewCircular(radius, height float64, segments int) (Column, error) {
	column := Column{Radius: radius, Height: height, Segments: segments, Circular: true}
	if _, err := column.Mesh(); err != nil {
		return Column{}, err
	}
	return column, nil
}

// Type is the .bim element type for a column.
func (c Column) Type() string { return "Column" }

// Info records the column dimensions in meters.
func (c Column) Info() map[string]string {
	if c.Circular {
		return map[string]string{
			"Section":  "Circular",
			"Radius":   format(c.Radius),
			"Height":   format(c.Height),
			"Segments": strconv.Itoa(c.Segments),
		}
	}
	return map[string]string{
		"Section": "Rectangular",
		"Width":   format(c.Width),
		"Depth":   format(c.Depth),
		"Height":  format(c.Height),
	}
}

// Mesh returns the column as a box or a capped cylinder.
func (c Column) Mesh() (*dotbim.Mesh, error) {
	if c.Circular {
		return primitive.Cylinder(c.Radius, c.Height, c.Segments, true)
	}
	return primitive.Box(c.Width, c.Depth, c.Height)
}

func format(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}
