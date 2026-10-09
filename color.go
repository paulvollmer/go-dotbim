package dotbim

// Color represents a color using red, green, blue, and alpha (transparency) values.
type Color struct {
	// The red component of the color.
	// Thrown if the value is not between 0 and 255.
	R uint8 `json:"r"`

	// The green component of the color.
	// Thrown if the value is not between 0 and 255.
	G uint8 `json:"g"`

	// The blue component of the color.
	// Thrown if the value is not between 0 and 255.
	B uint8 `json:"b"`

	// The alpha (transparency) component of the color.
	// Thrown if the value is not between 0 and 255.
	A uint8 `json:"a"`
}

// NewColor returns a color with the given red, green, blue, and alpha
// components. Each component is in the range 0 to 255. An alpha of 255 is
// fully opaque and an alpha of 0 is fully transparent.
func NewColor(r, g, b, a uint8) *Color {
	return &Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
