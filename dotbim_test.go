package dotbim_test

import (
	"encoding/json"
	"testing"
	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulvollmer/go-dotbim"
)

func TestNewDotBIM(t *testing.T) {
	info := map[string]string{"Author": "test"}

	bim := dotbim.NewDotBIM(info)

	require.NotNil(t, bim)
	assert.Equal(t, "1.0.0", dotbim.SchemaVersion)
	assert.Equal(t, dotbim.SchemaVersion, bim.SchemaVersion)
	assert.NotNil(t, bim.Meshes)
	assert.Empty(t, bim.Meshes)
	assert.NotNil(t, bim.Elements)
	assert.Empty(t, bim.Elements)
	assert.Equal(t, info, bim.Info)
}

func TestNewDotBIMNilInfo(t *testing.T) {
	bim := dotbim.NewDotBIM(nil)

	assert.Nil(t, bim.Info)
}

func TestAddMeshAndElement(t *testing.T) {
	bim := dotbim.NewDotBIM(nil)
	firstMesh := *dotbim.NewMesh(0, []float64{0, 0, 0}, []int{0, 1, 2})
	secondMesh := *dotbim.NewMesh(1, []float64{1, 1, 1}, []int{0})
	firstElement := *dotbim.NewElement(0, dotbim.Vector{}, dotbim.Rotation{}, uuid.Nil(), "Wall", dotbim.Color{}, nil)
	secondElement := *dotbim.NewElement(1, dotbim.Vector{X: 1}, dotbim.Rotation{W: 1}, uuid.Nil(), "Slab", dotbim.Color{}, nil)

	bim.AddMesh(firstMesh)
	bim.AddMesh(secondMesh)
	bim.AddElement(firstElement)
	bim.AddElement(secondElement)

	assert.Equal(t, []dotbim.Mesh{firstMesh, secondMesh}, bim.Meshes)
	assert.Equal(t, []dotbim.Element{firstElement, secondElement}, bim.Elements)
}

func TestDotBIMBuilder(t *testing.T) {
	info := map[string]string{"Author": "test"}
	meshes := []dotbim.Mesh{*dotbim.NewMesh(1, []float64{0}, []int{0})}
	elements := []dotbim.Element{
		*dotbim.NewElement(1, dotbim.Vector{Z: 2}, dotbim.Rotation{W: 1}, uuid.Nil(), "Wall", dotbim.Color{A: 255}, nil),
	}

	bim := dotbim.NewDotBIMBuilder().
		SetSchemaVersion("1.2.0").
		SetMeshes(meshes).
		SetElements(elements).
		SetInfo(info).
		Build()

	require.NotNil(t, bim)
	assert.Equal(t, "1.2.0", bim.SchemaVersion)
	assert.Equal(t, meshes, bim.Meshes)
	assert.Equal(t, elements, bim.Elements)
	assert.Equal(t, info, bim.Info)
}

func TestDotBIMJSON(t *testing.T) {
	guid := uuid.MustParse("12345678-1234-1234-1234-123456789abc")
	bim := dotbim.NewDotBIM(map[string]string{"Author": "test"})
	bim.AddMesh(*dotbim.NewMesh(0, []float64{0, 0, 0}, []int{0, 1, 2}))
	bim.AddElement(*dotbim.NewElement(
		0,
		*dotbim.NewVector(1, 2, 3),
		*dotbim.NewRotation(0, 0, 0, 1),
		guid,
		"Wall",
		*dotbim.NewColor(255, 210, 0, 255),
		map[string]string{"Material": "Concrete"},
	))

	data, err := json.Marshal(bim)
	require.NoError(t, err)
	assert.JSONEq(t, `{
		"schema_version": "1.0.0",
		"meshes": [
			{
				"mesh_id": 0,
				"coordinates": [0, 0, 0],
				"indices": [0, 1, 2]
			}
		],
		"elements": [
			{
				"mesh_id": 0,
				"vector": {"x": 1, "y": 2, "z": 3},
				"rotation": {"qx": 0, "qy": 0, "qz": 0, "qw": 1},
				"guid": "12345678-1234-1234-1234-123456789abc",
				"type": "Wall",
				"color": {"r": 255, "g": 210, "b": 0, "a": 255},
				"info": {"Material": "Concrete"}
			}
		],
		"info": {"Author": "test"}
	}`, string(data))

	var decoded dotbim.DotBIM
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, *bim, decoded)
}
