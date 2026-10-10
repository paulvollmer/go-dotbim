package dotbim_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulvollmer/go-dotbim"
)

func TestNewMesh(t *testing.T) {
	coordinates := []float64{0, 0, 0, 1, 0, 0, 0, 1, 0}
	indices := []int{0, 1, 2}

	mesh := dotbim.NewMesh(0, coordinates, indices)

	require.NotNil(t, mesh)
	assert.Equal(t, 0, mesh.MeshID)
	assert.Equal(t, coordinates, mesh.Coordinates)
	assert.Equal(t, indices, mesh.Indices)
}

func TestMeshBuilder(t *testing.T) {
	coordinates := []float64{0, 0, 0, 1, 0, 0, 0, 1, 0}
	indices := []int{0, 1, 2}

	mesh := dotbim.NewMeshBuilder().
		SetMeshID(3).
		SetCoordinates(coordinates).
		SetIndices(indices).
		Build()

	require.NotNil(t, mesh)
	assert.Equal(t, dotbim.NewMesh(3, coordinates, indices), mesh)
}

func TestMeshBuilderStartsEmpty(t *testing.T) {
	first := dotbim.NewMeshBuilder().SetMeshID(1).Build()
	second := dotbim.NewMeshBuilder().Build()

	assert.Equal(t, 1, first.MeshID)
	assert.Equal(t, &dotbim.Mesh{}, second)
}
