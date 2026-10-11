package primitive_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

func TestCylinder(t *testing.T) {
	mesh, err := primitive.Cylinder(1, 2, 4, true)
	require.NoError(t, err)

	assertValidMesh(t, mesh)
	assertBounds(t, mesh, 2, 2, 2)
	assertOutward(t, mesh, 1, 1, 1)

	_, err = primitive.Cylinder(1, 2, 2, true)
	assert.ErrorIs(t, err, primitive.ErrSegments)
	_, err = primitive.Cylinder(0, 2, 4, true)
	assert.ErrorIs(t, err, primitive.ErrDimension)
}
