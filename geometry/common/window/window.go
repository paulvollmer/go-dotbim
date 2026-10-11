// Package window builds a window frame and glass pane.
package window

import (
	"errors"
	"strconv"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

// ErrFrame means the frame leaves no opening for the glass.
var ErrFrame = errors.New("frame must leave room for glass")

// Window is a rectangular window. Width runs along Y, thickness along X, and
// height along Z. Frame is the border width on the width-height face.
type Window struct {
	Width     float64
	Height    float64
	Thickness float64
	Frame     float64
}

// New returns a window.
func New(width, height, thickness, frame float64) (Window, error) {
	window := Window{Width: width, Height: height, Thickness: thickness, Frame: frame}
	if _, _, err := window.Meshes(); err != nil {
		return Window{}, err
	}
	return window, nil
}

// Type is the .bim element type for a window.
func (w Window) Type() string { return "Window" }

// Info records the window dimensions in meters.
func (w Window) Info() map[string]string {
	return map[string]string{
		"Width":     format(w.Width),
		"Height":    format(w.Height),
		"Thickness": format(w.Thickness),
		"Frame":     format(w.Frame),
	}
}

// Meshes returns the frame and the glass pane. The glass is a quarter of the
// frame thickness and is centered in the thickness, inset by the frame width.
func (w Window) Meshes() (frame, glass *dotbim.Mesh, err error) {
	if _, err := primitive.Box(w.Thickness, w.Width, w.Height); err != nil {
		return nil, nil, err
	}
	if !(w.Frame > 0) || !(2*w.Frame < w.Width) || !(2*w.Frame < w.Height) {
		return nil, nil, ErrFrame
	}

	left, err := primitive.Box(w.Thickness, w.Frame, w.Height)
	if err != nil {
		return nil, nil, err
	}
	right, err := primitive.Box(w.Thickness, w.Frame, w.Height)
	if err != nil {
		return nil, nil, err
	}
	railWidth := w.Width - 2*w.Frame
	bottom, err := primitive.Box(w.Thickness, railWidth, w.Frame)
	if err != nil {
		return nil, nil, err
	}
	top, err := primitive.Box(w.Thickness, railWidth, w.Frame)
	if err != nil {
		return nil, nil, err
	}
	frame = dotbim.Merge(
		left,
		right.Translate(0, w.Width-w.Frame, 0),
		bottom.Translate(0, w.Frame, 0),
		top.Translate(0, w.Frame, w.Height-w.Frame),
	)

	glassThickness := w.Thickness / 4
	glass, err = primitive.Box(glassThickness, railWidth, w.Height-2*w.Frame)
	if err != nil {
		return nil, nil, err
	}
	glass = glass.Translate((w.Thickness-glassThickness)/2, w.Frame, w.Frame)
	return frame, glass, nil
}

func format(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}
