package dotbim

const SchemaVersion = "1.0.0"

type DotBIM struct {
	// The schema version of the BIM file.
	SchemaVersion string `json:"schema_version"`

	// The list of meshes in the file.
	Meshes []Mesh `json:"meshes"`

	// The list of elements in the file.
	Elements []Element `json:"elements"`

	// Additional information about the file.
	Info map[string]string `json:"info"`
}

func NewDotBIM(info map[string]string) *DotBIM {
	return &DotBIM{
		SchemaVersion: SchemaVersion,
		Meshes:        make([]Mesh, 0),
		Elements:      make([]Element, 0),
		Info:          info,
	}
}

func (bim *DotBIM) AddMesh(mesh Mesh) {
	bim.Meshes = append(bim.Meshes, mesh)
}

func (bim *DotBIM) AddElement(element Element) {
	bim.Elements = append(bim.Elements, element)
}
