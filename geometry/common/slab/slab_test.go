package slab_test

import (
	"testing"

	"github.com/paulvollmer/go-dotbim/geometry/common/slab"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlabMatchesBox(t *testing.T) {
	plate, err := slab.New(5, 4, 0.2)
	require.NoError(t, err)

	got, err := plate.Mesh()
	require.NoError(t, err)
	want, err := primitive.Box(5, 4, 0.2)
	require.NoError(t, err)

	assert.Equal(t, want, got)
	assert.Equal(t, "Slab", plate.Type())
	assert.Equal(t, "0.2", plate.Info()["Thickness"])
}
