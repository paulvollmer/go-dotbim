# single-wall

Writes one concrete wall to a [`.bim`](https://github.com/paireks/dotbim) file.

The wall is a box in meters, the `.bim` coordinate unit. `units` converts the thickness from centimeters before the mesh is built.

| Axis | Size | Meaning |
| --- | --- | --- |
| X | 0.20 m | thickness |
| Y | 4.00 m | length |
| Z | 2.00 m | height |

The element sits at the origin with the identity rotation `(0, 0, 0, 1)`. Its type is `Wall`, its info records `Material` as `Concrete`, and its color is opaque gold `(255, 210, 0, 255)`. The file info sets `Author` to `test`. Schema version is `1.0.0`.

## Run

From this directory:

```bash
go run .
```

That writes `single-wall.bim` next to `main.go`.

From the repository root:

```bash
make build
./examples/single-wall/single-wall
```

## What the program does

1. Creates an empty file with `dotbim.NewDotBIM`.
2. Builds mesh `0` from eight corners and twelve triangles, two per face.
3. Places one element on that mesh with `dotbim.NewElementBuilder`.
4. Marshals the file as indented JSON and writes `single-wall.bim`.

`single-wall.bim` contains that one mesh and one element. Open it in a viewer that reads the dotbim schema.
