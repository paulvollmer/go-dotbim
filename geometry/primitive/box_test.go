package primitive_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

func TestBoxMatchesSingleWall(t *testing.T) {
	mesh, err := primitive.Box(0.2, 4, 2)
	require.NoError(t, err)

	assert.Equal(t, dotbim.Coordinates{
		0, 0, 0, 0.2, 0, 0, 0.2, 4, 0, 0, 4, 0,
		0, 0, 2, 0.2, 0, 2, 0.2, 4, 2, 0, 4, 2,
	}, mesh.Coordinates)
	assert.Equal(t, dotbim.Indices{
		0, 1, 2, 0, 2, 3,
		4, 6, 5, 4, 7, 6,
		0, 4, 5, 0, 5, 1,
		1, 5, 6, 1, 6, 2,
		2, 6, 7, 2, 7, 3,
		3, 7, 4, 3, 4, 0,
	}, mesh.Indices)
	assert.Equal(t, 0, mesh.MeshID)
	assertBounds(t, mesh, 0.2, 4, 2)
}

func TestBoxRejectsNonPositiveSize(t *testing.T) {
	_, err := primitive.Box(0, 1, 1)
	assert.ErrorIs(t, err, primitive.ErrDimension)
	_, err = primitive.Box(1, -1, 1)
	assert.ErrorIs(t, err, primitive.ErrDimension)
}
