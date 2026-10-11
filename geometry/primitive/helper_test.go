package primitive_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulvollmer/go-dotbim"
)

func assertValidMesh(t *testing.T, mesh *dotbim.Mesh) {
	t.Helper()
	require.NotEmpty(t, mesh.Coordinates)
	require.NotEmpty(t, mesh.Indices)
	assert.Zero(t, len(mesh.Coordinates)%3)
	assert.Zero(t, len(mesh.Indices)%3)
	verts := len(mesh.Coordinates) / 3
	for _, idx := range mesh.Indices {
		assert.GreaterOrEqual(t, idx, 0)
		assert.Less(t, idx, verts)
	}
}

func assertBounds(t *testing.T, mesh *dotbim.Mesh, width, depth, height float64) {
	t.Helper()
	min, max := bounds(mesh)
	assert.InDelta(t, 0, min[0], 1e-9)
	assert.InDelta(t, 0, min[1], 1e-9)
	assert.InDelta(t, 0, min[2], 1e-9)
	assert.InDelta(t, width, max[0], 1e-9)
	assert.InDelta(t, depth, max[1], 1e-9)
	assert.InDelta(t, height, max[2], 1e-9)
}

func bounds(mesh *dotbim.Mesh) (min, max [3]float64) {
	min = [3]float64{math.Inf(1), math.Inf(1), math.Inf(1)}
	max = [3]float64{math.Inf(-1), math.Inf(-1), math.Inf(-1)}
	for i := 0; i < len(mesh.Coordinates); i += 3 {
		for k := 0; k < 3; k++ {
			v := mesh.Coordinates[i+k]
			min[k] = math.Min(min[k], v)
			max[k] = math.Max(max[k], v)
		}
	}
	return min, max
}

func assertOutward(t *testing.T, mesh *dotbim.Mesh, cx, cy, cz float64) {
	t.Helper()
	for i := 0; i < len(mesh.Indices); i += 3 {
		a := vertex(mesh, mesh.Indices[i])
		b := vertex(mesh, mesh.Indices[i+1])
		c := vertex(mesh, mesh.Indices[i+2])
		ab := [3]float64{b[0] - a[0], b[1] - a[1], b[2] - a[2]}
		ac := [3]float64{c[0] - a[0], c[1] - a[1], c[2] - a[2]}
		n := cross(ab, ac)
		mid := [3]float64{(a[0] + b[0] + c[0]) / 3, (a[1] + b[1] + c[1]) / 3, (a[2] + b[2] + c[2]) / 3}
		out := [3]float64{mid[0] - cx, mid[1] - cy, mid[2] - cz}
		assert.Greater(t, dot(n, out), 0.0)
	}
}

func centroidInBox(mesh *dotbim.Mesh, minX, maxX, minY, maxY, minZ, maxZ float64) bool {
	const eps = 1e-9
	for i := 0; i < len(mesh.Indices); i += 3 {
		a := vertex(mesh, mesh.Indices[i])
		b := vertex(mesh, mesh.Indices[i+1])
		c := vertex(mesh, mesh.Indices[i+2])
		x := (a[0] + b[0] + c[0]) / 3
		y := (a[1] + b[1] + c[1]) / 3
		z := (a[2] + b[2] + c[2]) / 3
		if x > minX+eps && x < maxX-eps && y > minY+eps && y < maxY-eps && z > minZ+eps && z < maxZ-eps {
			return true
		}
	}
	return false
}

func vertex(mesh *dotbim.Mesh, index int) [3]float64 {
	i := index * 3
	return [3]float64{mesh.Coordinates[i], mesh.Coordinates[i+1], mesh.Coordinates[i+2]}
}

func cross(a, b [3]float64) [3]float64 {
	return [3]float64{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

func dot(a, b [3]float64) float64 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}
