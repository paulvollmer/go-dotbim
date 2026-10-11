package main

import (
	"encoding/json"
	"fmt"
	"os"
	"uuid"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/units"
)

func main() {
	building := dotbim.NewDotBIM(map[string]string{
		"Author": "test",
	})

	// Define a wall geometry
	wallWidth := units.Centimeter(20)
	wallLength := units.Meter(4.0)
	wallHeight := units.Meter(2.0)
	coords := dotbim.NewCoordinates()
	// Floor corners
	coords.Vertex(0, 0, 0)
	coords.Vertex(wallWidth, 0, 0)
	coords.Vertex(wallWidth, wallLength, 0)
	coords.Vertex(0, wallLength, 0)
	// Ceiling corners
	coords.Vertex(0, 0, wallHeight)
	coords.Vertex(wallWidth, 0, wallHeight)
	coords.Vertex(wallWidth, wallLength, wallHeight)
	coords.Vertex(0, wallLength, wallHeight)
	wall := dotbim.NewMeshBuilder().
		SetMeshID(0).
		SetCoordinates(*coords).
		// Triangles (indices) for the 6 faces of the room (exterior walls/volume)
		SetIndices(
			*dotbim.NewIndices().
				// Floor
				Tri(0, 1, 2).Tri(0, 2, 3).
				// Ceiling
				Tri(4, 6, 5).Tri(4, 7, 6).
				// Front wall
				Tri(0, 4, 5).Tri(0, 5, 1).
				// Right wall
				Tri(1, 5, 6).Tri(1, 6, 2).
				// Back wall
				Tri(2, 6, 7).Tri(2, 7, 3).
				// Left wall
				Tri(3, 7, 4).Tri(3, 4, 0),
		).
		Build()
	building.AddMesh(*wall)

	roomElement := dotbim.NewElementBuilder().
		SetMeshID(wall.MeshID).
		SetVector(*dotbim.NewVector(0, 0, 0)).
		SetRotation(*dotbim.NewRotation(0, 0, 0, 1)).
		SetGuid(uuid.New()).
		SetType("Wall").
		SetColor(*dotbim.NewColor(255, 210, 0, 255)).
		SetInfo(map[string]string{
			"Material": "Concrete",
		}).
		Build()
	building.AddElement(*roomElement)

	// Export as a JSON/.bim file
	writeBIM(building)
}

func writeBIM(building *dotbim.DotBIM) {
	// 3. Export as a JSON/.bim file
	fileData, err := json.MarshalIndent(building, "", "  ")
	if err != nil {
		fmt.Println("Fehler beim Generieren von JSON:", err)
		return
	}

	filename := "single-wall.bim"
	err = os.WriteFile(filename, fileData, 0644)
	if err != nil {
		fmt.Println("Fehler beim Schreiben der Datei:", err)
		return
	}

	fmt.Printf("Successfully written %q file\n", filename)
}
