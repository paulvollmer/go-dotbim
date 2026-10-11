package primitive_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

func TestExtrudeConvexMatchesBoxBounds(t *testing.T) {
	mesh, err := primitive.Extrude([]float64{
		0, 0,
		2, 0,
		2, 3,
		0, 3,
	}, 4, nil)
	require.NoError(t, err)

	assertValidMesh(t, mesh)
	assertBounds(t, mesh, 2, 3, 4)
	assert.Len(t, mesh.Indices, 36)
}

func TestExtrudeHole(t *testing.T) {
	mesh, err := primitive.Extrude([]float64{
		0, 0,
		4, 0,
		4, 4,
		0, 4,
	}, 1, []primitive.Rect{{X: 1, Y: 1, Width: 1, Height: 1}})
	require.NoError(t, err)

	assertValidMesh(t, mesh)
	assertBounds(t, mesh, 4, 4, 1)
	assert.False(t, centroidInBox(mesh, 0, 1, 1, 2, 1, 2))

	_, err = primitive.Extrude([]float64{0, 0, 1, 0, 1, 1, 0, 1}, 1, []primitive.Rect{{X: 0, Y: 0, Width: 0.5, Height: 0.5}})
	assert.ErrorIs(t, err, primitive.ErrHole)
	_, err = primitive.Extrude([]float64{0, 0, 0, 1, 1, 0}, 1, nil)
	assert.ErrorIs(t, err, primitive.ErrProfile)
}
