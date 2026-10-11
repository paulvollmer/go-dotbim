# single-wall

Writes one concrete wall to a [`.bim`](https://github.com/paireks/dotbim) file. The program builds a solid box by hand: eight corners with `Coordinates.Vertex` and twelve triangles with `Indices.Tri`, two per face, then places one element on that mesh.

Lengths are meters, the `.bim` coordinate unit. Thickness is `units.Centimeter(20)`. Length is `units.Meter(4.0)` and height is `units.Meter(2.0)`. `units.Meter` returns its value unchanged.

| Axis | Size | Meaning |
| --- | --- | --- |
| X | 0.20 m | thickness |
| Y | 4.00 m | length |
| Z | 2.00 m | height |

The element sits at the origin. Its rotation is the identity `(x, y, z, w) = (0, 0, 0, 1)` from `dotbim.NewRotation`. Its type is `Wall`, its info records `Material` as `Concrete`, and its color is opaque gold RGBA `(255, 210, 0, 255)` from `dotbim.NewColor`. The file info sets `Author` to `test`. Schema version is `1.0.0`.

## Run

From this directory:

```bash
go run .
```

That writes `single-wall.bim` next to `main.go` and prints `Successfully written "single-wall.bim" file`. Each run calls `uuid.New()`, so the element `guid` changes. The file is listed in `.gitignore`.

From the repository root:

```bash
make build
./examples/single-wall/single-wall
```

Drag `single-wall.bim` onto [Online 3D Viewer](https://3dviewer.net/). Other apps that read the format are in the [dotbim README](https://github.com/paireks/dotbim#apps-supporting-bim).

## What the program does

1. Builds an in-memory file with `dotbim.NewDotBIM`. The schema version is `1.0.0` and `Author` is `test`.
2. Builds mesh `0` from eight corners and twelve triangles, two per face.
3. Places one element on that mesh with `dotbim.NewElementBuilder`. The element `guid` comes from `uuid.New()`.
4. Marshals the value as indented JSON and writes `single-wall.bim`.

The file contains that one mesh and one element.

## See also

- [Repository README](../../README.md)
- [examples/all-primitives](../all-primitives)
- [examples/simple-house](../simple-house)
