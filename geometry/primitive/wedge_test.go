package primitive_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

func TestWedge(t *testing.T) {
	mesh, err := primitive.Wedge(2, 3, 1)
	require.NoError(t, err)

	assertValidMesh(t, mesh)
	assertBounds(t, mesh, 2, 3, 1)
	_, err = primitive.Wedge(1, 0, 1)
	assert.ErrorIs(t, err, primitive.ErrDimension)
}
