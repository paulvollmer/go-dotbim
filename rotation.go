package dotbim

// Represents a rotation in three-dimensional space using a quaternion.
type Rotation struct {
	// The x-coordinate of the quaternion.
	X float64 `json:"qx"`

	// The y-coordinate of the quaternion.
	Y float64 `json:"qy"`

	// The z-coordinate of the quaternion.
	Z float64 `json:"qz"`

	// The w-coordinate of the quaternion.
	W float64 `json:"qw"`
}

// NewRotation returns a quaternion rotation with components x, y, z, and w.
// The identity rotation, which applies no turn, is (0, 0, 0, 1).
func NewRotation(x, y, z, w float64) *Rotation {
	return &Rotation{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}
