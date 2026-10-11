package dotbim

// Represents a mesh object in three-dimensional space.
//
//go:generate go tool github.com/paulvollmer/go-genbuilder
type Mesh struct {
	// The identifier for the mesh.
	// The value of MeshId should be greater than or equal to 0.
	MeshID int `json:"mesh_id"`

	// The list of coordinates for the mesh vertices.
	Coordinates Coordinates `json:"coordinates"`

	// The list of indices for the mesh triangles.
	Indices Indices `json:"indices"`
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

// Coordinates lists vertex positions as consecutive x, y, z values.
type Coordinates []float64

// NewCoordinates returns an empty coordinate list.
func NewCoordinates() *Coordinates {
	return &Coordinates{}
}

// Vertex appends one vertex and returns its index.
func (c *Coordinates) Vertex(x, y, z float64) int {
	*c = append(*c, x, y, z)
	return len(*c)/3 - 1
}

// Indices lists triangle corners in groups of three.
type Indices []int

// NewIndices returns an empty triangle list.
func NewIndices() *Indices {
	return &Indices{}
}

// Tri appends one triangle and returns idx so calls can chain.
func (idx *Indices) Tri(i, j, k int) *Indices {
	*idx = append(*idx, i, j, k)
	return idx
}

// Translate returns a copy of the mesh moved by x, y, z.
// The receiver is left unchanged.
func (m *Mesh) Translate(x, y, z float64) *Mesh {
	out := &Mesh{
		MeshID:      m.MeshID,
		Coordinates: append(Coordinates(nil), m.Coordinates...),
		Indices:     append(Indices(nil), m.Indices...),
	}
	for i := 0; i < len(out.Coordinates); i += 3 {
		out.Coordinates[i] += x
		out.Coordinates[i+1] += y
		out.Coordinates[i+2] += z
	}
	return out
}

// Merge concatenates meshes into one. Vertices are not welded.
// Each part must be non-nil. The result uses mesh id 0.
func Merge(parts ...*Mesh) *Mesh {
	var coords Coordinates
	var indices Indices
	for _, part := range parts {
		base := len(coords) / 3
		coords = append(coords, part.Coordinates...)
		for _, id := range part.Indices {
			indices = append(indices, id+base)
		}
	}
	return &Mesh{Coordinates: coords, Indices: indices}
}
