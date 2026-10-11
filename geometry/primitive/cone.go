package primitive

import (
	"math"

	"github.com/paulvollmer/go-dotbim"
)

// Cone returns a cone. The base is on z = 0 and the apex is at z = height.
// The base is centered on (radius, radius) so the bounding box starts at the
// origin. segments must be at least 3. capped adds the base disc.
// The mesh id is 0.
func Cone(radius, height float64, segments int, capped bool) (*dotbim.Mesh, error) {
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
	base := make([]int, segments)
	for i := 0; i < segments; i++ {
		theta := 2 * math.Pi * float64(i) / float64(segments)
		x := radius + radius*math.Cos(theta)
		y := radius + radius*math.Sin(theta)
		base[i] = coords.Vertex(x, y, 0)
	}
	apex := coords.Vertex(radius, radius, height)

	for i := 0; i < segments; i++ {
		next := (i + 1) % segments
		indices.Tri(base[i], base[next], apex)
	}

	if capped {
		center := coords.Vertex(radius, radius, 0)
		for i := 0; i < segments; i++ {
			next := (i + 1) % segments
			indices.Tri(center, base[next], base[i])
		}
	}

	return mesh(coords, indices), nil
}
