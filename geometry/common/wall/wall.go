// Package wall builds a straight wall mesh from geometry primitives.
package wall

import (
	"errors"
	"strconv"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

// ErrOpening means the opening is not strictly inside the wall face.
var ErrOpening = errors.New("opening must lie inside the wall")

// Opening is a rectangular hole through the wall thickness.
// Offset is the distance along the wall from its minimum corner.
// Sill is the distance from the base to the bottom of the hole.
type Opening struct {
	Offset float64
	Sill   float64
	Width  float64
	Height float64
}

// Wall is a straight wall. Length runs along Y, thickness along X, and height
// along Z.
type Wall struct {
	Length    float64
	Height    float64
	Thickness float64
	Opening   *Opening
}

// New returns a wall. opening is optional.
func New(length, height, thickness float64, opening *Opening) (Wall, error) {
	wall := Wall{Length: length, Height: height, Thickness: thickness}
	if opening != nil {
		copied := *opening
		wall.Opening = &copied
	}
	if _, err := wall.Mesh(); err != nil {
		return Wall{}, err
	}
	return wall, nil
}

// Type is the .bim element type for a wall.
func (w Wall) Type() string { return "Wall" }

// Info records the wall dimensions in meters.
func (w Wall) Info() map[string]string {
	info := map[string]string{
		"Length":    format(w.Length),
		"Height":    format(w.Height),
		"Thickness": format(w.Thickness),
	}
	if w.Opening != nil {
		info["OpeningOffset"] = format(w.Opening.Offset)
		info["OpeningSill"] = format(w.Opening.Sill)
		info["OpeningWidth"] = format(w.Opening.Width)
		info["OpeningHeight"] = format(w.Opening.Height)
	}
	return info
}

// Mesh returns the wall mesh. A wall with an opening is an extrusion of the
// length-height profile through the thickness.
func (w Wall) Mesh() (*dotbim.Mesh, error) {
	if w.Opening == nil {
		return primitive.Box(w.Thickness, w.Length, w.Height)
	}
	if err := w.validateOpening(); err != nil {
		return nil, err
	}
	mesh, err := primitive.Extrude([]float64{
		0, 0,
		w.Length, 0,
		w.Length, w.Height,
		0, w.Height,
	}, w.Thickness, []primitive.Rect{{
		X:      w.Opening.Offset,
		Y:      w.Opening.Sill,
		Width:  w.Opening.Width,
		Height: w.Opening.Height,
	}})
	if err != nil {
		return nil, err
	}
	return swapThickness(mesh), nil
}

func (w Wall) validateOpening() error {
	if _, err := primitive.Box(w.Thickness, w.Length, w.Height); err != nil {
		return err
	}
	opening := w.Opening
	if opening.Width <= 0 || opening.Height <= 0 ||
		!(opening.Offset > 0) || !(opening.Sill > 0) ||
		!(opening.Offset+opening.Width < w.Length) ||
		!(opening.Sill+opening.Height < w.Height) {
		return ErrOpening
	}
	return nil
}

// swapThickness turns an extrusion of the length-height profile into wall axes.
// The extrusion uses X for length, Y for height, and Z for thickness.
// The copy uses X for thickness, Y for length, and Z for height.
func swapThickness(mesh *dotbim.Mesh) *dotbim.Mesh {
	out := &dotbim.Mesh{
		Coordinates: append(dotbim.Coordinates(nil), mesh.Coordinates...),
		Indices:     append(dotbim.Indices(nil), mesh.Indices...),
	}
	for i := 0; i < len(out.Coordinates); i += 3 {
		x, y, z := out.Coordinates[i], out.Coordinates[i+1], out.Coordinates[i+2]
		out.Coordinates[i] = z
		out.Coordinates[i+1] = x
		out.Coordinates[i+2] = y
	}
	return out
}

func format(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}
