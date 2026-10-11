package primitive

import "github.com/paulvollmer/go-dotbim"

// Wedge returns a right-angle ramp. The face at y = 0 has the full height, and
// the height is zero at y = depth. The mesh id is 0.
func Wedge(width, depth, height float64) (*dotbim.Mesh, error) {
	if err := dimension("width", width); err != nil {
		return nil, err
	}
	if err := dimension("depth", depth); err != nil {
		return nil, err
	}
	if err := dimension("height", height); err != nil {
		return nil, err
	}

	coords := dotbim.NewCoordinates()
	coords.Vertex(0, 0, 0)
	coords.Vertex(width, 0, 0)
	coords.Vertex(width, depth, 0)
	coords.Vertex(0, depth, 0)
	coords.Vertex(0, 0, height)
	coords.Vertex(width, 0, height)

	indices := dotbim.NewIndices().
		Tri(0, 1, 2).
		Tri(0, 2, 3).
		Tri(0, 4, 5).
		Tri(0, 5, 1).
		Tri(3, 4, 0).
		Tri(1, 5, 2).
		Tri(4, 5, 2).
		Tri(4, 2, 3)

	return mesh(coords, indices), nil
}
