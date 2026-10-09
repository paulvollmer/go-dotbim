package dotbim

// Represents a mesh object in three-dimensional space.
//
//go:generate go tool github.com/paulvollmer/go-genbuilder
type Mesh struct {
	// The identifier for the mesh.
	// The value of MeshId should be greater than or equal to 0.
	MeshID int `json:"mesh_id"`

	// The list of coordinates for the mesh vertices.
	Coordinates []float64 `json:"coordinates"`

	// The list of indices for the mesh triangles.
	Indices []int `json:"indices"`
}

// NewMesh returns a mesh with the given identifier, vertices, and triangles.
// meshID should be greater than or equal to 0. coordinates lists vertex
// positions as consecutive x, y, z values. indices lists triangle corners in
// groups of three, each value an index into those vertices.
func NewMesh(meshID int, coordinates []float64, indices []int) *Mesh {
	return &Mesh{
		MeshID:      meshID,
		Coordinates: coordinates,
		Indices:     indices,
	}
}
