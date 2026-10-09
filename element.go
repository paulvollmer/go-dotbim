package dotbim

import (
	"uuid"
)

// Represents an element in a BIM file.
type Element struct {
	// The identifier of the mesh associated with the element.
	MeshID int `json:"mesh_id"`

	// The position of the element.
	Vector Vector `json:"vector"`

	// The rotation of the element.
	Rotation Rotation `json:"rotation"`

	// The globally unique identifier for the element.
	// Thrown if the value is not a valid globally unique identifier.
	Guid uuid.UUID `json:"guid"`

	// The type of the element.
	Type string `json:"type"`

	// The color of the element.
	Color Color `json:"color"`

	// Additional information about the element.
	Info map[string]string `json:"info"`
}

func NewElement(meshID int, vector Vector, rotation Rotation, guid uuid.UUID, typ string, color Color, info map[string]string) *Element {
	return &Element{
		MeshID:   meshID,
		Vector:   vector,
		Rotation: rotation,
		Guid:     guid,
		Type:     typ,
		Color:    color,
		Info:     info,
	}
}
