package door_test

import (
	"testing"

	"github.com/paulvollmer/go-dotbim/geometry/common/door"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDoorMatchesBox(t *testing.T) {
	leaf, err := door.New(0.9, 2.1, 0.04)
	require.NoError(t, err)

	got, err := leaf.Mesh()
	require.NoError(t, err)
	want, err := primitive.Box(0.04, 0.9, 2.1)
	require.NoError(t, err)

	assert.Equal(t, want, got)
	assert.Equal(t, "Door", leaf.Type())
	assert.Equal(t, map[string]string{
		"Width":     "0.9",
		"Height":    "2.1",
		"Thickness": "0.04",
	}, leaf.Info())
}

func TestDoorRejectsNonPositiveSize(t *testing.T) {
	_, err := door.New(0.9, 0, 0.04)
	assert.ErrorIs(t, err, primitive.ErrDimension)
}
