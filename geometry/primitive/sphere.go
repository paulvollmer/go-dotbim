package primitive

import (
	"math"

	"github.com/paulvollmer/go-dotbim"
)

// Sphere returns a latitude-longitude sphere. The sphere is centered on
// (radius, radius, radius) so its bounding box starts at the origin.
// segments is the number of longitude sides and must be at least 3. rings is
// the number of stacks from pole to pole and must be at least 2.
// The mesh id is 0.
func Sphere(radius float64, segments, rings int) (*dotbim.Mesh, error) {
	if err := dimension("radius", radius); err != nil {
		return nil, err
	}
	if segments < 3 {
		return nil, ErrSegments
	}
	if rings < 2 {
		return nil, ErrRings
	}

	coords := dotbim.NewCoordinates()
	indices := dotbim.NewIndices()
	cx, cy, cz := radius, radius, radius
	north := coords.Vertex(cx, cy, cz+radius)
	ringsV := make([][]int, rings-1)
	for y := 1; y < rings; y++ {
		phi := math.Pi * float64(y) / float64(rings)
		row := make([]int, segments)
		for x := 0; x < segments; x++ {
			theta := 2 * math.Pi * float64(x) / float64(segments)
			row[x] = coords.Vertex(
				cx+radius*math.Sin(phi)*math.Cos(theta),
				cy+radius*math.Sin(phi)*math.Sin(theta),
				cz+radius*math.Cos(phi),
			)
		}
		ringsV[y-1] = row
	}
	south := coords.Vertex(cx, cy, cz-radius)

	first := ringsV[0]
	for x := 0; x < segments; x++ {
		next := (x + 1) % segments
		indices.Tri(north, first[x], first[next])
	}
	for y := 0; y < len(ringsV)-1; y++ {
		upper := ringsV[y]
		lower := ringsV[y+1]
		for x := 0; x < segments; x++ {
			next := (x + 1) % segments
			indices.Tri(upper[x], lower[x], lower[next])
			indices.Tri(upper[x], lower[next], upper[next])
		}
	}
	last := ringsV[len(ringsV)-1]
	for x := 0; x < segments; x++ {
		next := (x + 1) % segments
		indices.Tri(south, last[next], last[x])
	}

	return mesh(coords, indices), nil
}
