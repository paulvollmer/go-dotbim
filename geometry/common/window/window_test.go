package window_test

import (
	"math"
	"testing"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/common/window"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWindowFrameAndGlass(t *testing.T) {
	win, err := window.New(1.2, 1.4, 0.08, 0.05)
	require.NoError(t, err)

	frame, glass, err := win.Meshes()
	require.NoError(t, err)
	assertValid(t, frame)
	assertValid(t, glass)
	assertBounds(t, frame, 0.08, 1.2, 1.4)

	glassMin, glassMax := bounds(glass)
	assert.Greater(t, glassMin[0], 0.0)
	assert.Less(t, glassMax[0], 0.08)
	assert.Greater(t, glassMin[1], 0.0)
	assert.Less(t, glassMax[1], 1.2)
	assert.Greater(t, glassMin[2], 0.0)
	assert.Less(t, glassMax[2], 1.4)
	assert.Equal(t, "Window", win.Type())
	assert.Equal(t, "0.05", win.Info()["Frame"])
}

func TestWindowRejectsFrame(t *testing.T) {
	_, err := window.New(1, 1, 0.1, 0.5)
	assert.ErrorIs(t, err, window.ErrFrame)
	_, err = window.New(0, 1, 0.1, 0.05)
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
			min[k] = math.Min(min[k], mesh.Coordinates[i+k])
			max[k] = math.Max(max[k], mesh.Coordinates[i+k])
		}
	}
	return min, max
}
