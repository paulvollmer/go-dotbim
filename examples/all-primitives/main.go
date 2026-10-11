// Command all-primitives writes every geometry primitive to one .bim file.
// Each shape is one element, placed in a row along X on the ground plane.
// Constructors return mesh id 0. This program sets the id before adding the mesh.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"uuid"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/primitive"
)

const (
	size     = 1.0
	radius   = size / 2
	segments = 32
	rings    = 16
	gap      = 0.5
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	file := dotbim.NewDotBIM(map[string]string{
		"Author":  "example",
		"Example": "all-primitives",
	})
	identity := *dotbim.NewRotation(0, 0, 0, 1)
	x := 0.0

	box, err := primitive.Box(size, size, size)
	if err != nil {
		return err
	}
	box.MeshID = 0
	file.AddMesh(*box)
	file.AddElement(*dotbim.NewElement(
		box.MeshID,
		*dotbim.NewVector(x, 0, 0),
		identity,
		uuid.New(),
		"Box",
		*dotbim.NewColor(190, 190, 190, 255),
		map[string]string{
			"Width":  "1",
			"Depth":  "1",
			"Height": "1",
		},
	))
	x += size + gap

	wedge, err := primitive.Wedge(size, size, size)
	if err != nil {
		return err
	}
	wedge.MeshID = 1
	file.AddMesh(*wedge)
	file.AddElement(*dotbim.NewElement(
		wedge.MeshID,
		*dotbim.NewVector(x, 0, 0),
		identity,
		uuid.New(),
		"Wedge",
		*dotbim.NewColor(214, 122, 56, 255),
		map[string]string{
			"Width":  "1",
			"Depth":  "1",
			"Height": "1",
		},
	))
	x += size + gap

	cylinder, err := primitive.Cylinder(radius, size, segments, true)
	if err != nil {
		return err
	}
	cylinder.MeshID = 2
	file.AddMesh(*cylinder)
	file.AddElement(*dotbim.NewElement(
		cylinder.MeshID,
		*dotbim.NewVector(x, 0, 0),
		identity,
		uuid.New(),
		"Cylinder",
		*dotbim.NewColor(66, 135, 201, 255),
		map[string]string{
			"Radius":   "0.5",
			"Height":   "1",
			"Segments": "32",
			"Capped":   "true",
		},
	))
	x += size + gap

	cone, err := primitive.Cone(radius, size, segments, true)
	if err != nil {
		return err
	}
	cone.MeshID = 3
	file.AddMesh(*cone)
	file.AddElement(*dotbim.NewElement(
		cone.MeshID,
		*dotbim.NewVector(x, 0, 0),
		identity,
		uuid.New(),
		"Cone",
		*dotbim.NewColor(196, 72, 72, 255),
		map[string]string{
			"Radius":   "0.5",
			"Height":   "1",
			"Segments": "32",
			"Capped":   "true",
		},
	))
	x += size + gap

	sphere, err := primitive.Sphere(radius, segments, rings)
	if err != nil {
		return err
	}
	sphere.MeshID = 4
	file.AddMesh(*sphere)
	file.AddElement(*dotbim.NewElement(
		sphere.MeshID,
		*dotbim.NewVector(x, 0, 0),
		identity,
		uuid.New(),
		"Sphere",
		*dotbim.NewColor(72, 166, 104, 255),
		map[string]string{
			"Radius":   "0.5",
			"Segments": "32",
			"Rings":    "16",
		},
	))
	x += size + gap

	triangle, err := primitive.Extrude([]float64{
		0, 0,
		size, 0,
		size / 2, size,
	}, size, nil)
	if err != nil {
		return err
	}
	triangle.MeshID = 5
	file.AddMesh(*triangle)
	file.AddElement(*dotbim.NewElement(
		triangle.MeshID,
		*dotbim.NewVector(x, 0, 0),
		identity,
		uuid.New(),
		"Extrude",
		*dotbim.NewColor(150, 96, 196, 255),
		map[string]string{
			"Profile": "triangle",
			"Height":  "1",
		},
	))
	x += size + gap

	holed, err := primitive.Extrude([]float64{
		0, 0,
		size, 0,
		size, size,
		0, size,
	}, size, []primitive.Rect{{
		X: 0.25, Y: 0.25, Width: 0.5, Height: 0.5,
	}})
	if err != nil {
		return err
	}
	holed.MeshID = 6
	file.AddMesh(*holed)
	file.AddElement(*dotbim.NewElement(
		holed.MeshID,
		*dotbim.NewVector(x, 0, 0),
		identity,
		uuid.New(),
		"Extrude",
		*dotbim.NewColor(212, 176, 55, 255),
		map[string]string{
			"Profile":    "rectangle",
			"Height":     "1",
			"HoleX":      "0.25",
			"HoleY":      "0.25",
			"HoleWidth":  "0.5",
			"HoleHeight": "0.5",
		},
	))

	return writeBIM(file)
}

func writeBIM(file *dotbim.DotBIM) error {
	data, err := json.MarshalIndent(file, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal .bim: %w", err)
	}
	const filename = "all-primitives.bim"
	if err := os.WriteFile(filename, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", filename, err)
	}
	fmt.Printf("Successfully written %q file\n", filename)
	return nil
}
