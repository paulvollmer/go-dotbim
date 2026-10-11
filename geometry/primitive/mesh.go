// Package primitive builds triangulated meshes in meters for .bim coordinates.
//
// Every shape returns a *dotbim.Mesh with mesh id 0. The caller sets the id
// when the mesh is added to a file. The origin of every shape is its minimum
// corner. X is width, Y is depth, and Z is up. Callers convert other units
// before calling these functions.
package primitive

import (
	"errors"
	"fmt"

	"github.com/paulvollmer/go-dotbim"
)

var (
	// ErrDimension means a size was not greater than zero.
	ErrDimension = errors.New("dimension must be greater than zero")

	// ErrSegments means a revolution was given fewer than 3 segments.
	ErrSegments = errors.New("segments must be at least 3")

	// ErrRings means a sphere was given fewer than 2 rings.
	ErrRings = errors.New("rings must be at least 2")

	// ErrProfile means the outline cannot be extruded.
	ErrProfile = errors.New("profile must be a counter-clockwise polygon")

	// ErrHole means a hole is outside the profile or overlaps another hole.
	ErrHole = errors.New("hole must lie inside the profile")
)

func dimension(name string, value float64) error {
	if !(value > 0) {
		return fmt.Errorf("%w: %s", ErrDimension, name)
	}
	return nil
}

func mesh(coordinates *dotbim.Coordinates, indices *dotbim.Indices) *dotbim.Mesh {
	return &dotbim.Mesh{Coordinates: *coordinates, Indices: *indices}
}
