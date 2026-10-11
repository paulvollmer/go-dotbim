package primitive

import "github.com/paulvollmer/go-dotbim"

// Box returns a rectangular box. width is along X, depth along Y, and height
// along Z. The corner order matches the single-wall example. The mesh id is 0.
func Box(width, depth, height float64) (*dotbim.Mesh, error) {
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
	coords.Vertex(width, depth, height)
	coords.Vertex(0, depth, height)

	indices := dotbim.NewIndices().
		Tri(0, 1, 2).
		Tri(0, 2, 3).
		Tri(4, 6, 5).
		Tri(4, 7, 6).
		Tri(0, 4, 5).
		Tri(0, 5, 1).
		Tri(1, 5, 6).
		Tri(1, 6, 2).
		Tri(2, 6, 7).
		Tri(2, 7, 3).
		Tri(3, 7, 4).
		Tri(3, 4, 0)

	return mesh(coords, indices), nil
}
