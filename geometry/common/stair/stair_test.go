package stair_test

import (
	"math"
	"testing"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/common/stair"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStair(t *testing.T) {
	flight, err := stair.New(3, 1, 2, 4)
	require.NoError(t, err)

	mesh, err := flight.Mesh()
	require.NoError(t, err)
	require.NotEmpty(t, mesh.Indices)
	assert.Zero(t, len(mesh.Coordinates)%3)
	assert.Zero(t, len(mesh.Indices)%3)
	assert.Len(t, mesh.Coordinates, 4*8*3)

	min, max := bounds(mesh)
	assert.InDelta(t, 0, min[0], 1e-9)
	assert.InDelta(t, 0, min[1], 1e-9)
	assert.InDelta(t, 0, min[2], 1e-9)
	assert.InDelta(t, 1, max[0], 1e-9)
	assert.InDelta(t, 3, max[1], 1e-9)
	assert.InDelta(t, 2, max[2], 1e-9)
	assert.Equal(t, "Stair", flight.Type())
	assert.Equal(t, "4", flight.Info()["Steps"])
}

func TestStairRejectsStepCount(t *testing.T) {
	_, err := stair.New(3, 1, 2, 0)
	assert.ErrorIs(t, err, stair.ErrSteps)
	_, err = stair.New(0, 1, 2, 2)
	assert.ErrorIs(t, err, primitive.ErrDimension)
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
