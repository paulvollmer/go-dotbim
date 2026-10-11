package primitive_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

func TestCone(t *testing.T) {
	mesh, err := primitive.Cone(1, 3, 4, true)
	require.NoError(t, err)

	assertValidMesh(t, mesh)
	assertBounds(t, mesh, 2, 2, 3)
	assertOutward(t, mesh, 1, 1, 1)

	open, err := primitive.Cone(1, 3, 4, false)
	require.NoError(t, err)
	assertValidMesh(t, open)
	assert.Less(t, len(open.Indices), len(mesh.Indices))
}
