package primitive_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

func TestSphere(t *testing.T) {
	mesh, err := primitive.Sphere(1, 4, 4)
	require.NoError(t, err)

	assertValidMesh(t, mesh)
	assertBounds(t, mesh, 2, 2, 2)
	assertOutward(t, mesh, 1, 1, 1)

	_, err = primitive.Sphere(1, 2, 4)
	assert.ErrorIs(t, err, primitive.ErrSegments)
	_, err = primitive.Sphere(1, 4, 1)
	assert.ErrorIs(t, err, primitive.ErrRings)
}
