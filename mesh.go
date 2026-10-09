package dotbim

// Represents a mesh object in three-dimensional space.
type Mesh struct {
	// The identifier for the mesh.
	// The value of MeshId should be greater than or equal to 0.
	MeshID int `json:"mesh_id"`

	// The list of coordinates for the mesh vertices.
	Coordinates []float64 `json:"coordinates"`

	// The list of indices for the mesh triangles.
	Indices []int `json:"indices"`
}

func NewMesh(meshID int, coordinates []float64, indices []int) *Mesh {
	return &Mesh{
		MeshID:      meshID,
		Coordinates: coordinates,
		Indices:     indices,
	}
}
