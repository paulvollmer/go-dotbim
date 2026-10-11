package column_test

import (
	"testing"

	"github.com/paulvollmer/go-dotbim/geometry/common/column"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRectangularColumnMatchesBox(t *testing.T) {
	col, err := column.New(0.3, 0.4, 3)
	require.NoError(t, err)

	got, err := col.Mesh()
	require.NoError(t, err)
	want, err := primitive.Box(0.3, 0.4, 3)
	require.NoError(t, err)

	assert.Equal(t, want, got)
	assert.Equal(t, "Column", col.Type())
	assert.Equal(t, "Rectangular", col.Info()["Section"])
}

func TestCircularColumnMatchesCylinder(t *testing.T) {
	col, err := column.NewCircular(0.2, 3, 12)
	require.NoError(t, err)

	got, err := col.Mesh()
	require.NoError(t, err)
	want, err := primitive.Cylinder(0.2, 3, 12, true)
	require.NoError(t, err)

	assert.Equal(t, want, got)
	assert.Equal(t, "Circular", col.Info()["Section"])
	assert.Equal(t, "12", col.Info()["Segments"])
}

func TestColumnRejectsBadSize(t *testing.T) {
	_, err := column.NewCircular(0.2, 3, 2)
	assert.ErrorIs(t, err, primitive.ErrSegments)
}
