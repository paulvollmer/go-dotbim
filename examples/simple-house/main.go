// Command simple-house writes a house built from every geometry/common type.
//
// X is width, Y is depth, and Z is up. Each common mesh keeps its own axes:
// length or width along Y, thickness along X, and height along Z. A wall,
// door, or beam whose length runs along world X is turned with the element
// rotation. The element vector is applied after that rotation.
package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"uuid"

	"github.com/paulvollmer/go-dotbim"
	"github.com/paulvollmer/go-dotbim/geometry/common/beam"
	"github.com/paulvollmer/go-dotbim/geometry/common/column"
	"github.com/paulvollmer/go-dotbim/geometry/common/door"
	"github.com/paulvollmer/go-dotbim/geometry/common/roof"
	"github.com/paulvollmer/go-dotbim/geometry/common/slab"
	"github.com/paulvollmer/go-dotbim/geometry/common/stair"
	"github.com/paulvollmer/go-dotbim/geometry/common/wall"
	"github.com/paulvollmer/go-dotbim/geometry/common/window"
)

const (
	wallT  = 0.20
	wallH  = 2.60
	houseW = 8.0
	houseD = 6.0
	porchD = 1.80
	floorT = 0.90
	stairL = 2.00
	stairW = 1.20
	steps  = 6

	doorW    = 0.90
	doorH    = 2.10
	doorT    = 0.05
	doorSill = 0.05

	winW     = 1.40
	winH     = 1.20
	winT     = 0.08
	winFrame = 0.07
	winSill  = 1.00

	colR   = 0.16
	colSeg = 24
	beamW  = 0.20
	beamD  = 0.30

	roofRise = 2.0
	overhang = 0.40

	chimneyW = 0.55
	chimneyD = 0.55
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
		"Example": "simple-house",
	})

	quarter := math.Sqrt2 / 2
	identity := *dotbim.NewRotation(0, 0, 0, 1)
	// +90° about Z: local +Y goes to world -X, and local +X goes to world +Y.
	faceX := *dotbim.NewRotation(0, 0, quarter, quarter)
	// -90° about Z: local +Y goes to world +X, and local +X goes to world -Y.
	spanX := *dotbim.NewRotation(0, 0, -quarter, quarter)

	houseX := 0.50
	houseY := stairL + porchD
	floorZ := floorT
	sideLen := houseD - 2*wallT
	sideY := houseY + wallT
	roofY := stairL + 0.15
	roofDepth := houseY + houseD + overhang - roofY
	roofZ := floorZ + wallH

	floor, err := slab.New(houseW, porchD+houseD, floorT)
	if err != nil {
		return err
	}
	floorMesh, err := floor.Mesh()
	if err != nil {
		return err
	}
	floorMesh.MeshID = 0
	file.AddMesh(*floorMesh)
	floorInfo := floor.Info()
	floorInfo["Name"] = "Floor"
	file.AddElement(*dotbim.NewElement(
		floorMesh.MeshID,
		*dotbim.NewVector(houseX, stairL, 0),
		identity,
		uuid.New(),
		floor.Type(),
		*dotbim.NewColor(176, 174, 168, 255),
		floorInfo,
	))

	front, err := wall.New(houseW, wallH, wallT, &wall.Opening{
		Offset: (houseW - doorW) / 2,
		Sill:   doorSill,
		Width:  doorW,
		Height: doorH,
	})
	if err != nil {
		return err
	}
	frontMesh, err := front.Mesh()
	if err != nil {
		return err
	}
	frontMesh.MeshID = 1
	file.AddMesh(*frontMesh)
	frontInfo := front.Info()
	frontInfo["Name"] = "Front"
	// After faceX the wall extends toward -X, so the vector is its maximum X.
	file.AddElement(*dotbim.NewElement(
		frontMesh.MeshID,
		*dotbim.NewVector(houseX+houseW, houseY, floorZ),
		faceX,
		uuid.New(),
		front.Type(),
		*dotbim.NewColor(236, 230, 216, 255),
		frontInfo,
	))

	back, err := wall.New(houseW, wallH, wallT, nil)
	if err != nil {
		return err
	}
	backMesh, err := back.Mesh()
	if err != nil {
		return err
	}
	backMesh.MeshID = 2
	file.AddMesh(*backMesh)
	backInfo := back.Info()
	backInfo["Name"] = "Back"
	file.AddElement(*dotbim.NewElement(
		backMesh.MeshID,
		*dotbim.NewVector(houseX+houseW, houseY+houseD-wallT, floorZ),
		faceX,
		uuid.New(),
		back.Type(),
		*dotbim.NewColor(236, 230, 216, 255),
		backInfo,
	))

	winOffset := (sideLen - winW) / 2
	sideOpening := &wall.Opening{Offset: winOffset, Sill: winSill, Width: winW, Height: winH}
	left, err := wall.New(sideLen, wallH, wallT, sideOpening)
	if err != nil {
		return err
	}
	leftMesh, err := left.Mesh()
	if err != nil {
		return err
	}
	leftMesh.MeshID = 3
	file.AddMesh(*leftMesh)
	leftInfo := left.Info()
	leftInfo["Name"] = "Left"
	file.AddElement(*dotbim.NewElement(
		leftMesh.MeshID,
		*dotbim.NewVector(houseX, sideY, floorZ),
		identity,
		uuid.New(),
		left.Type(),
		*dotbim.NewColor(236, 230, 216, 255),
		leftInfo,
	))

	right, err := wall.New(sideLen, wallH, wallT, sideOpening)
	if err != nil {
		return err
	}
	rightMesh, err := right.Mesh()
	if err != nil {
		return err
	}
	rightMesh.MeshID = 4
	file.AddMesh(*rightMesh)
	rightInfo := right.Info()
	rightInfo["Name"] = "Right"
	file.AddElement(*dotbim.NewElement(
		rightMesh.MeshID,
		*dotbim.NewVector(houseX+houseW-wallT, sideY, floorZ),
		identity,
		uuid.New(),
		right.Type(),
		*dotbim.NewColor(236, 230, 216, 255),
		rightInfo,
	))

	entrance, err := door.New(doorW, doorH, doorT)
	if err != nil {
		return err
	}
	doorMesh, err := entrance.Mesh()
	if err != nil {
		return err
	}
	doorMesh.MeshID = 5
	file.AddMesh(*doorMesh)
	doorInfo := entrance.Info()
	doorInfo["Name"] = "Entrance"
	doorX := houseX + (houseW-doorW)/2
	doorY := houseY + (wallT-doorT)/2
	file.AddElement(*dotbim.NewElement(
		doorMesh.MeshID,
		*dotbim.NewVector(doorX+doorW, doorY, floorZ+doorSill),
		faceX,
		uuid.New(),
		entrance.Type(),
		*dotbim.NewColor(122, 74, 44, 255),
		doorInfo,
	))

	opening, err := window.New(winW, winH, winT, winFrame)
	if err != nil {
		return err
	}
	frameMesh, glassMesh, err := opening.Meshes()
	if err != nil {
		return err
	}
	frameMesh.MeshID = 6
	glassMesh.MeshID = 7
	file.AddMesh(*frameMesh)
	file.AddMesh(*glassMesh)
	winXInset := (wallT - winT) / 2
	leftFrame := opening.Info()
	leftFrame["Name"] = "Left"
	leftFrame["Part"] = "Frame"
	file.AddElement(*dotbim.NewElement(
		frameMesh.MeshID,
		*dotbim.NewVector(houseX+winXInset, sideY+winOffset, floorZ+winSill),
		identity,
		uuid.New(),
		opening.Type(),
		*dotbim.NewColor(248, 248, 246, 255),
		leftFrame,
	))
	leftGlass := opening.Info()
	leftGlass["Name"] = "Left"
	leftGlass["Part"] = "Glass"
	file.AddElement(*dotbim.NewElement(
		glassMesh.MeshID,
		*dotbim.NewVector(houseX+winXInset, sideY+winOffset, floorZ+winSill),
		identity,
		uuid.New(),
		opening.Type(),
		*dotbim.NewColor(170, 206, 220, 96),
		leftGlass,
	))
	rightFrame := opening.Info()
	rightFrame["Name"] = "Right"
	rightFrame["Part"] = "Frame"
	file.AddElement(*dotbim.NewElement(
		frameMesh.MeshID,
		*dotbim.NewVector(houseX+houseW-wallT+winXInset, sideY+winOffset, floorZ+winSill),
		identity,
		uuid.New(),
		opening.Type(),
		*dotbim.NewColor(248, 248, 246, 255),
		rightFrame,
	))
	rightGlass := opening.Info()
	rightGlass["Name"] = "Right"
	rightGlass["Part"] = "Glass"
	file.AddElement(*dotbim.NewElement(
		glassMesh.MeshID,
		*dotbim.NewVector(houseX+houseW-wallT+winXInset, sideY+winOffset, floorZ+winSill),
		identity,
		uuid.New(),
		opening.Type(),
		*dotbim.NewColor(170, 206, 220, 96),
		rightGlass,
	))

	colH := wallH - beamD
	porchCol, err := column.NewCircular(colR, colH, colSeg)
	if err != nil {
		return err
	}
	colMesh, err := porchCol.Mesh()
	if err != nil {
		return err
	}
	colMesh.MeshID = 8
	file.AddMesh(*colMesh)
	leftColX := houseX + 0.60
	rightColX := houseX + houseW - 0.60
	colY := stairL + 0.55
	leftColInfo := porchCol.Info()
	leftColInfo["Name"] = "Porch left"
	file.AddElement(*dotbim.NewElement(
		colMesh.MeshID,
		*dotbim.NewVector(leftColX-colR, colY-colR, floorZ),
		identity,
		uuid.New(),
		porchCol.Type(),
		*dotbim.NewColor(154, 152, 146, 255),
		leftColInfo,
	))
	rightColInfo := porchCol.Info()
	rightColInfo["Name"] = "Porch right"
	file.AddElement(*dotbim.NewElement(
		colMesh.MeshID,
		*dotbim.NewVector(rightColX-colR, colY-colR, floorZ),
		identity,
		uuid.New(),
		porchCol.Type(),
		*dotbim.NewColor(154, 152, 146, 255),
		rightColInfo,
	))

	header, err := beam.New(beamW, rightColX-leftColX+2*colR, beamD)
	if err != nil {
		return err
	}
	beamMesh, err := header.Mesh()
	if err != nil {
		return err
	}
	beamMesh.MeshID = 9
	file.AddMesh(*beamMesh)
	beamInfo := header.Info()
	beamInfo["Name"] = "Porch"
	// After spanX the beam width extends toward -Y, so the vector is its maximum Y.
	file.AddElement(*dotbim.NewElement(
		beamMesh.MeshID,
		*dotbim.NewVector(leftColX-colR, colY+beamW/2, floorZ+colH),
		spanX,
		uuid.New(),
		header.Type(),
		*dotbim.NewColor(96, 62, 40, 255),
		beamInfo,
	))

	top, err := roof.New(houseW+2*overhang, roofDepth, roofRise)
	if err != nil {
		return err
	}
	roofMesh, err := top.Mesh()
	if err != nil {
		return err
	}
	roofMesh.MeshID = 10
	file.AddMesh(*roofMesh)
	roofInfo := top.Info()
	roofInfo["Name"] = "Roof"
	file.AddElement(*dotbim.NewElement(
		roofMesh.MeshID,
		*dotbim.NewVector(houseX-overhang, roofY, roofZ),
		identity,
		uuid.New(),
		top.Type(),
		*dotbim.NewColor(176, 78, 58, 255),
		roofInfo,
	))

	flight, err := stair.New(stairL, stairW, floorT, steps)
	if err != nil {
		return err
	}
	stairMesh, err := flight.Mesh()
	if err != nil {
		return err
	}
	stairMesh.MeshID = 11
	file.AddMesh(*stairMesh)
	stairInfo := flight.Info()
	stairInfo["Name"] = "Entrance"
	file.AddElement(*dotbim.NewElement(
		stairMesh.MeshID,
		*dotbim.NewVector(houseX+(houseW-stairW)/2, 0, 0),
		identity,
		uuid.New(),
		flight.Type(),
		*dotbim.NewColor(164, 162, 156, 255),
		stairInfo,
	))

	chimneyX := houseX + houseW - wallT - 0.25 - chimneyW
	chimneyY := houseY + houseD - wallT - 0.25 - chimneyD
	along := (chimneyY + chimneyD/2) - roofY
	t := along / roofDepth
	if t < 0 {
		t = 0
	}
	if t > 1 {
		t = 1
	}
	chimneyTop := roofZ + roofRise*(1-t) + 0.70
	stack, err := column.New(chimneyW, chimneyD, chimneyTop-floorZ)
	if err != nil {
		return err
	}
	chimneyMesh, err := stack.Mesh()
	if err != nil {
		return err
	}
	chimneyMesh.MeshID = 12
	file.AddMesh(*chimneyMesh)
	chimneyInfo := stack.Info()
	chimneyInfo["Name"] = "Chimney"
	file.AddElement(*dotbim.NewElement(
		chimneyMesh.MeshID,
		*dotbim.NewVector(chimneyX, chimneyY, floorZ),
		identity,
		uuid.New(),
		stack.Type(),
		*dotbim.NewColor(148, 72, 58, 255),
		chimneyInfo,
	))

	return writeBIM(file)
}

func writeBIM(file *dotbim.DotBIM) error {
	data, err := json.MarshalIndent(file, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal .bim: %w", err)
	}
	const filename = "simple-house.bim"
	if err := os.WriteFile(filename, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", filename, err)
	}
	fmt.Printf("Successfully written %q file\n", filename)
	return nil
}
