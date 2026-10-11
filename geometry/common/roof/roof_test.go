package roof_test

import (
	"testing"

	"github.com/paulvollmer/go-dotbim/geometry/common/roof"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoofMatchesWedge(t *testing.T) {
	slope, err := roof.New(6, 4, 2)
	require.NoError(t, err)

	got, err := slope.Mesh()
	require.NoError(t, err)
	want, err := primitive.Wedge(6, 4, 2)
	require.NoError(t, err)

	assert.Equal(t, want, got)
	assert.Equal(t, "Roof", slope.Type())
	assert.Equal(t, "2", slope.Info()["Rise"])
}
