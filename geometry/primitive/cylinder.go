package primitive

import (
	"math"

	"github.com/paulvollmer/go-dotbim"
)

// Cylinder returns a cylinder of the given radius and height. The axis is Z
// and the base sits on z = 0. The bounding box starts at the origin, so the
// axis passes through (radius, radius). segments is the number of sides and
// must be at least 3. capped adds the top and bottom discs. The mesh id is 0.
func Cylinder(radius, height float64, segments int, capped bool) (*dotbim.Mesh, error) {
	if err := dimension("radius", radius); err != nil {
		return nil, err
	}
	if err := dimension("height", height); err != nil {
		return nil, err
	}
	if segments < 3 {
		return nil, ErrSegments
	}

	coords := dotbim.NewCoordinates()
	indices := dotbim.NewIndices()
	bottom := make([]int, segments)
	top := make([]int, segments)
	for i := 0; i < segments; i++ {
		theta := 2 * math.Pi * float64(i) / float64(segments)
		x := radius + radius*math.Cos(theta)
		y := radius + radius*math.Sin(theta)
		bottom[i] = coords.Vertex(x, y, 0)
		top[i] = coords.Vertex(x, y, height)
	}

	for i := 0; i < segments; i++ {
		next := (i + 1) % segments
		indices.Tri(bottom[i], bottom[next], top[next])
		indices.Tri(bottom[i], top[next], top[i])
	}

	if capped {
		base := coords.Vertex(radius, radius, 0)
		crown := coords.Vertex(radius, radius, height)
		for i := 0; i < segments; i++ {
			next := (i + 1) % segments
			indices.Tri(base, bottom[next], bottom[i])
			indices.Tri(crown, top[i], top[next])
		}
	}

	return mesh(coords, indices), nil
}
