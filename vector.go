package dotbim

// Vector represents a three-dimensional vector with double-precision floating-point coordinates.
type Vector struct {
	// The x-coordinate of the vector.
	X float64 `json:"x"`

	// The y-coordinate of the vector.
	Y float64 `json:"y"`

	// The z-coordinate of the vector.
	Z float64 `json:"z"`
}

// NewVector returns a three-dimensional vector with the given coordinates.
func NewVector(x, y, z float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
	}
}
