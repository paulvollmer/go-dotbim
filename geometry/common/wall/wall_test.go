package wall_test

import (
	"math"
	"testing"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/common/wall"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolidWallMatchesBox(t *testing.T) {
	gotWall, err := wall.New(4, 2, 0.2, nil)
	require.NoError(t, err)

	got, err := gotWall.Mesh()
	require.NoError(t, err)
	want, err := primitive.Box(0.2, 4, 2)
	require.NoError(t, err)

	assert.Equal(t, want, got)
	assert.Equal(t, "Wall", gotWall.Type())
	assert.Equal(t, "4", gotWall.Info()["Length"])
	assert.Equal(t, "2", gotWall.Info()["Height"])
	assert.Equal(t, "0.2", gotWall.Info()["Thickness"])
}

func TestWallOpeningHasNoTriangleInside(t *testing.T) {
	opening := &wall.Opening{Offset: 1, Sill: 0.5, Width: 1, Height: 1}
	gotWall, err := wall.New(4, 2, 0.2, opening)
	require.NoError(t, err)

	mesh, err := gotWall.Mesh()
	require.NoError(t, err)
	assertValid(t, mesh)
	assertBounds(t, mesh, 0.2, 4, 2)
	assert.Greater(t, len(mesh.Coordinates)/3, 8)
	assert.False(t, centroidInside(mesh, 0, 0.2, 1, 2, 0.5, 1.5))
	assert.Equal(t, "1", gotWall.Info()["OpeningOffset"])
	assert.Equal(t, "0.5", gotWall.Info()["OpeningSill"])
}

func TestWallRejectsOpeningOnTheBoundary(t *testing.T) {
	_, err := wall.New(4, 2, 0.2, &wall.Opening{Offset: 0, Sill: 0.5, Width: 1, Height: 1})
	assert.ErrorIs(t, err, wall.ErrOpening)
	_, err = wall.New(0, 2, 0.2, nil)
	assert.ErrorIs(t, err, primitive.ErrDimension)
}

func assertValid(t *testing.T, mesh *dotbim.Mesh) {
	t.Helper()
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
	min := [3]float64{math.Inf(1), math.Inf(1), math.Inf(1)}
	max := [3]float64{math.Inf(-1), math.Inf(-1), math.Inf(-1)}
	for i := 0; i < len(mesh.Coordinates); i += 3 {
		for k := 0; k < 3; k++ {
			min[k] = math.Min(min[k], mesh.Coordinates[i+k])
			max[k] = math.Max(max[k], mesh.Coordinates[i+k])
		}
	}
	assert.InDelta(t, 0, min[0], 1e-9)
	assert.InDelta(t, 0, min[1], 1e-9)
	assert.InDelta(t, 0, min[2], 1e-9)
	assert.InDelta(t, width, max[0], 1e-9)
	assert.InDelta(t, depth, max[1], 1e-9)
	assert.InDelta(t, height, max[2], 1e-9)
}

func centroidInside(mesh *dotbim.Mesh, minX, maxX, minY, maxY, minZ, maxZ float64) bool {
	const eps = 1e-9
	for i := 0; i < len(mesh.Indices); i += 3 {
		var c [3]float64
		for _, idx := range mesh.Indices[i : i+3] {
			c[0] += mesh.Coordinates[idx*3]
			c[1] += mesh.Coordinates[idx*3+1]
			c[2] += mesh.Coordinates[idx*3+2]
		}
		c[0] /= 3
		c[1] /= 3
		c[2] /= 3
		if c[0] > minX+eps && c[0] < maxX-eps && c[1] > minY+eps && c[1] < maxY-eps && c[2] > minZ+eps && c[2] < maxZ-eps {
			return true
		}
	}
	return false
}
