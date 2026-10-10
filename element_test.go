package dotbim_test

import (
	"testing"
	"uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulvollmer/go-dotbim"
)

func TestNewElement(t *testing.T) {
	guid := uuid.MustParse("12345678-1234-1234-1234-123456789abc")
	vector := *dotbim.NewVector(1, 2, 3)
	rotation := *dotbim.NewRotation(0, 0, 0, 1)
	color := *dotbim.NewColor(255, 210, 0, 255)
	info := map[string]string{"Material": "Concrete"}

	element := dotbim.NewElement(7, vector, rotation, guid, "Wall", color, info)

	require.NotNil(t, element)
	assert.Equal(t, 7, element.MeshID)
	assert.Equal(t, vector, element.Vector)
	assert.Equal(t, rotation, element.Rotation)
	assert.Equal(t, guid, element.Guid)
	assert.Equal(t, "Wall", element.Type)
	assert.Equal(t, color, element.Color)
	assert.Equal(t, info, element.Info)
}

func TestElementBuilder(t *testing.T) {
	guid := uuid.MustParse("12345678-1234-1234-1234-123456789abc")
	vector := *dotbim.NewVector(1, 2, 3)
	rotation := *dotbim.NewRotation(0, 0, 0, 1)
	color := *dotbim.NewColor(255, 210, 0, 255)
	info := map[string]string{"Material": "Concrete"}

	element := dotbim.NewElementBuilder().
		SetMeshID(7).
		SetVector(vector).
		SetRotation(rotation).
		SetGuid(guid).
		SetType("Wall").
		SetColor(color).
		SetInfo(info).
		Build()

	require.NotNil(t, element)
	assert.Equal(t, dotbim.NewElement(7, vector, rotation, guid, "Wall", color, info), element)
}

func TestElementBuilderStartsEmpty(t *testing.T) {
	element := dotbim.NewElementBuilder().Build()

	assert.Equal(t, &dotbim.Element{}, element)
}
