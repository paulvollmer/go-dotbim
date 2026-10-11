package beam_test

import (
	"testing"

	"github.com/paulvollmer/go-dotbim/geometry/common/beam"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBeamMatchesBox(t *testing.T) {
	member, err := beam.New(0.2, 4, 0.4)
	require.NoError(t, err)

	got, err := member.Mesh()
	require.NoError(t, err)
	want, err := primitive.Box(0.2, 4, 0.4)
	require.NoError(t, err)

	assert.Equal(t, want, got)
	assert.Equal(t, "Beam", member.Type())
	assert.Equal(t, "4", member.Info()["Length"])
}
