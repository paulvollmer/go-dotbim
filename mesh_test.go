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
	assert.Equal(t, dotbim.Coordinates(coordinates), mesh.Coordinates)
	assert.Equal(t, dotbim.Indices(indices), mesh.Indices)
}

func TestVertexIndexAndTriChain(t *testing.T) {
	coords := dotbim.NewCoordinates()
	assert.Equal(t, 0, coords.Vertex(1, 2, 3))
	assert.Equal(t, 1, coords.Vertex(4, 5, 6))
	assert.Equal(t, dotbim.Coordinates{1, 2, 3, 4, 5, 6}, *coords)

	indices := dotbim.NewIndices().Tri(0, 1, 2).Tri(2, 1, 3)
	assert.Equal(t, dotbim.Indices{0, 1, 2, 2, 1, 3}, *indices)
}

func TestTranslateLeavesTheOriginal(t *testing.T) {
	mesh := dotbim.NewMesh(3, []float64{0, 0, 0, 1, 0, 0}, []int{0, 1, 0})

	moved := mesh.Translate(1, 2, 3)

	assert.Equal(t, dotbim.Coordinates{1, 2, 3, 2, 2, 3}, moved.Coordinates)
	assert.Equal(t, dotbim.Indices{0, 1, 0}, moved.Indices)
	assert.Equal(t, 3, moved.MeshID)
	assert.Equal(t, dotbim.Coordinates{0, 0, 0, 1, 0, 0}, mesh.Coordinates)
}

func TestMergeOffsetsIndices(t *testing.T) {
	a := dotbim.NewMesh(0, []float64{0, 0, 0}, []int{0, 0, 0})
	b := dotbim.NewMesh(1, []float64{1, 0, 0}, []int{0, 0, 0})

	got := dotbim.Merge(a, b)

	assert.Equal(t, 0, got.MeshID)
	assert.Equal(t, dotbim.Coordinates{0, 0, 0, 1, 0, 0}, got.Coordinates)
	assert.Equal(t, dotbim.Indices{0, 0, 0, 1, 1, 1}, got.Indices)
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
