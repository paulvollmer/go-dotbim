# all-primitives

Writes every mesh from `geometry/primitive` to one [`.bim`](https://github.com/paireks/dotbim) file.

Each shape is 1 m in X, Y, and Z. Round shapes use a radius of 0.5 m, so the bounding box is still 1 m. Coordinates are meters. The elements sit on `z = 0` in a row along X, with a 0.5 m gap and no rotation.

| Element | Shape | Color |
| --- | --- | --- |
| Box | 1 × 1 × 1 m box | gray `(190, 190, 190)` |
| Wedge | ramp, full height at `y = 0` | orange `(214, 122, 56)` |
| Cylinder | 32 sides, capped | blue `(66, 135, 201)` |
| Cone | 32 sides, capped | red `(196, 72, 72)` |
| Sphere | 32 segments, 16 rings | green `(72, 166, 104)` |
| Extrude | triangle profile, height 1 m | purple `(150, 96, 196)` |
| Extrude | 1 × 1 m rectangle with a 0.5 m square hole | gold `(212, 176, 55)` |

`Extrude` is one constructor. The file includes it twice: a convex profile, and a rectangular profile with a hole.

## Run

From this directory:

```bash
go run .
```

That writes `all-primitives.bim` next to `main.go`.

From the repository root:

```bash
make build
./examples/all-primitives/all-primitives
```

Open `all-primitives.bim` in a viewer that reads the dotbim schema.
