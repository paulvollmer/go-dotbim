package primitive

import (
	"sort"

	"github.com/paulvollmer/go-dotbim"
)

// Rect is an axis-aligned rectangle in the XY profile of an extrusion.
type Rect struct {
	X, Y          float64
	Width, Height float64
}

type point struct {
	x, y float64
}

// Extrude builds a vertical solid from a closed XY profile. profile is
// consecutive x, y values, counter-clockwise when viewed from above, without
// repeating the first point. height is the extrusion along Z.
//
// Without holes the profile must be convex. Rectangular holes are supported
// when the profile itself is an axis-aligned rectangle. Each hole must lie
// strictly inside that rectangle and must not overlap another hole.
func Extrude(profile []float64, height float64, holes []Rect) (*dotbim.Mesh, error) {
	if err := dimension("height", height); err != nil {
		return nil, err
	}
	pts, err := profilePoints(profile)
	if err != nil {
		return nil, err
	}
	if len(holes) == 0 {
		if !convexCCW(pts) {
			return nil, ErrProfile
		}
		return extrudeConvex(pts, height), nil
	}
	minX, minY, maxX, maxY, ok := rectangleBounds(pts)
	if !ok {
		return nil, ErrProfile
	}
	if err := validateHoles(minX, minY, maxX, maxY, holes); err != nil {
		return nil, err
	}
	return extrudeRectangle(minX, minY, maxX, maxY, height, holes), nil
}

func profilePoints(profile []float64) ([]point, error) {
	if len(profile) < 6 || len(profile)%2 != 0 {
		return nil, ErrProfile
	}
	pts := make([]point, 0, len(profile)/2)
	for i := 0; i < len(profile); i += 2 {
		pts = append(pts, point{profile[i], profile[i+1]})
	}
	if pts[0] == pts[len(pts)-1] {
		pts = pts[:len(pts)-1]
	}
	if len(pts) < 3 {
		return nil, ErrProfile
	}
	return pts, nil
}

func convexCCW(pts []point) bool {
	for i := range pts {
		a := pts[i]
		b := pts[(i+1)%len(pts)]
		c := pts[(i+2)%len(pts)]
		cross := (b.x-a.x)*(c.y-b.y) - (b.y-a.y)*(c.x-b.x)
		if !(cross > 0) {
			return false
		}
	}
	return true
}

func rectangleBounds(pts []point) (minX, minY, maxX, maxY float64, ok bool) {
	if len(pts) != 4 {
		return 0, 0, 0, 0, false
	}
	minX, minY = pts[0].x, pts[0].y
	maxX, maxY = minX, minY
	for _, p := range pts[1:] {
		minX = min(minX, p.x)
		minY = min(minY, p.y)
		maxX = max(maxX, p.x)
		maxY = max(maxY, p.y)
	}
	if !(maxX > minX) || !(maxY > minY) {
		return 0, 0, 0, 0, false
	}
	corners := map[[2]float64]struct{}{}
	for _, p := range pts {
		onX := p.x == minX || p.x == maxX
		onY := p.y == minY || p.y == maxY
		if !onX || !onY {
			return 0, 0, 0, 0, false
		}
		corners[[2]float64{p.x, p.y}] = struct{}{}
	}
	if len(corners) != 4 || !convexCCW(pts) {
		return 0, 0, 0, 0, false
	}
	return minX, minY, maxX, maxY, true
}

func validateHoles(minX, minY, maxX, maxY float64, holes []Rect) error {
	for i, hole := range holes {
		if err := dimension("hole width", hole.Width); err != nil {
			return err
		}
		if err := dimension("hole height", hole.Height); err != nil {
			return err
		}
		if !(hole.X > minX) || !(hole.Y > minY) || !(hole.X+hole.Width < maxX) || !(hole.Y+hole.Height < maxY) {
			return ErrHole
		}
		for _, other := range holes[i+1:] {
			if hole.X < other.X+other.Width && other.X < hole.X+hole.Width &&
				hole.Y < other.Y+other.Height && other.Y < hole.Y+hole.Height {
				return ErrHole
			}
		}
	}
	return nil
}

func extrudeConvex(pts []point, height float64) *dotbim.Mesh {
	coords := dotbim.NewCoordinates()
	indices := dotbim.NewIndices()
	n := len(pts)
	bottom := make([]int, n)
	top := make([]int, n)
	for i, p := range pts {
		bottom[i] = coords.Vertex(p.x, p.y, 0)
		top[i] = coords.Vertex(p.x, p.y, height)
	}
	for i := 0; i < n; i++ {
		next := (i + 1) % n
		indices.Tri(bottom[i], bottom[next], top[next])
		indices.Tri(bottom[i], top[next], top[i])
	}
	for i := 1; i < n-1; i++ {
		indices.Tri(bottom[0], bottom[i], bottom[i+1])
		indices.Tri(top[0], top[i+1], top[i])
	}
	return mesh(coords, indices)
}

func extrudeRectangle(minX, minY, maxX, maxY, height float64, holes []Rect) *dotbim.Mesh {
	xs := []float64{minX, maxX}
	ys := []float64{minY, maxY}
	for _, hole := range holes {
		xs = append(xs, hole.X, hole.X+hole.Width)
		ys = append(ys, hole.Y, hole.Y+hole.Height)
	}
	sort.Float64s(xs)
	sort.Float64s(ys)
	xs = uniqueFloats(xs)
	ys = uniqueFloats(ys)

	coords := dotbim.NewCoordinates()
	indices := dotbim.NewIndices()
	for i := 0; i < len(xs)-1; i++ {
		for j := 0; j < len(ys)-1; j++ {
			x0, x1 := xs[i], xs[i+1]
			y0, y1 := ys[j], ys[j+1]
			if !(x1 > x0) || !(y1 > y0) {
				continue
			}
			if pointInHole((x0+x1)/2, (y0+y1)/2, holes) {
				continue
			}
			addCap(coords, indices, x0, y0, x1, y1, 0, false)
			addCap(coords, indices, x0, y0, x1, y1, height, true)
		}
	}
	addLoop(coords, indices, []point{{minX, minY}, {maxX, minY}, {maxX, maxY}, {minX, maxY}}, height)
	for _, hole := range holes {
		addLoop(coords, indices, []point{
			{hole.X, hole.Y},
			{hole.X + hole.Width, hole.Y},
			{hole.X + hole.Width, hole.Y + hole.Height},
			{hole.X, hole.Y + hole.Height},
		}, height)
	}
	return mesh(coords, indices)
}

func uniqueFloats(values []float64) []float64 {
	out := values[:0]
	var prev float64
	for i, v := range values {
		if i == 0 || v != prev {
			out = append(out, v)
			prev = v
		}
	}
	return out
}

func pointInHole(x, y float64, holes []Rect) bool {
	for _, hole := range holes {
		if x > hole.X && x < hole.X+hole.Width && y > hole.Y && y < hole.Y+hole.Height {
			return true
		}
	}
	return false
}

func addCap(coords *dotbim.Coordinates, indices *dotbim.Indices, x0, y0, x1, y1, z float64, top bool) {
	p00 := coords.Vertex(x0, y0, z)
	p10 := coords.Vertex(x1, y0, z)
	p11 := coords.Vertex(x1, y1, z)
	p01 := coords.Vertex(x0, y1, z)
	if top {
		indices.Tri(p00, p01, p11)
		indices.Tri(p00, p11, p10)
		return
	}
	indices.Tri(p00, p10, p11)
	indices.Tri(p00, p11, p01)
}

func addLoop(coords *dotbim.Coordinates, indices *dotbim.Indices, pts []point, height float64) {
	for i := range pts {
		next := pts[(i+1)%len(pts)]
		cur := pts[i]
		b0 := coords.Vertex(cur.x, cur.y, 0)
		b1 := coords.Vertex(next.x, next.y, 0)
		t1 := coords.Vertex(next.x, next.y, height)
		t0 := coords.Vertex(cur.x, cur.y, height)
		indices.Tri(b0, b1, t1)
		indices.Tri(b0, t1, t0)
	}
}
