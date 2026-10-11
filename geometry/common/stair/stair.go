// Package stair builds a straight flight of steps.
package stair

import (
	"errors"
	"strconv"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

// ErrSteps means the flight has fewer than one step.
var ErrSteps = errors.New("step count must be at least 1")

// Stair is a straight flight. Length runs along Y, width along X, and height
// along Z. The steps divide the length and height equally.
type Stair struct {
	Length float64
	Width  float64
	Height float64
	Steps  int
}

// New returns a stair flight.
func New(length, width, height float64, steps int) (Stair, error) {
	stair := Stair{Length: length, Width: width, Height: height, Steps: steps}
	if _, err := stair.Mesh(); err != nil {
		return Stair{}, err
	}
	return stair, nil
}

// Type is the .bim element type for a stair.
func (s Stair) Type() string { return "Stair" }

// Info records the stair dimensions in meters.
func (s Stair) Info() map[string]string {
	return map[string]string{
		"Length": format(s.Length),
		"Width":  format(s.Width),
		"Height": format(s.Height),
		"Steps":  strconv.Itoa(s.Steps),
	}
}

// Mesh returns one merged mesh, with a box for each tread.
func (s Stair) Mesh() (*dotbim.Mesh, error) {
	if s.Steps < 1 {
		return nil, ErrSteps
	}
	rise := s.Height / float64(s.Steps)
	run := s.Length / float64(s.Steps)
	parts := make([]*dotbim.Mesh, 0, s.Steps)
	for i := 0; i < s.Steps; i++ {
		step, err := primitive.Box(s.Width, run, float64(i+1)*rise)
		if err != nil {
			return nil, err
		}
		parts = append(parts, step.Translate(0, float64(i)*run, 0))
	}
	return dotbim.Merge(parts...), nil
}

func format(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}
