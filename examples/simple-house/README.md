# simple-house

Writes a house to one [`.bim`](https://github.com/paireks/dotbim) file. The model uses every type in `geometry/common`.

X is width, Y is depth, and Z is up. Lengths are meters. You walk up the stair onto a raised floor, cross a porch, and enter through the front door. The roof is a single slope, high over the porch and low at the back.

Common meshes keep length along Y and thickness along X. The front wall, the back wall, the door, and the porch beam turn a quarter turn about Z with the element rotation, so their length runs along world X. The element vector is applied after that rotation. The two porch posts share one column mesh, and the two windows share one frame mesh and one glass mesh.

```
y
^
|  +----------------------+  back wall
|  |                  [c] |  chimney
|  |                      |
|  | [window]    [window] |
|  +--------[door]--------+  front wall
|  | (o) beam         (o) |  porch columns
|  +----------------------+  floor slab
|           stair
+----------------------------> x
```

| Part | Geometry | Where it sits |
| --- | --- | --- |
| Floor | Slab | 8 × 7.8 m, 0.90 m thick, porch and house |
| Front, back, left, right | Wall | 0.20 m thick, 2.60 m high. Front has the door opening. Each side wall has a window opening. |
| Entrance | Door | 0.90 × 2.10 m leaf, centered in the front wall |
| Left, right | Window | Frame and glass in each side wall. Sill at 1.00 m. |
| Porch left, porch right | Column | Circular posts, radius 0.16 m |
| Chimney | Column | Rectangular stack through the roof |
| Porch | Beam | Header resting on the posts and meeting the roof |
| Roof | Roof | Shed slope, 0.40 m overhang, 2.00 m rise at the porch |
| Entrance | Stair | 6 steps, 1.20 m wide, rising 0.90 m to the floor |

A wall opening has to sit strictly inside the face, so the door sill is 5 cm above the floor.

## Run

From this directory:

```bash
go run .
```

That writes `simple-house.bim` next to `main.go`.

From the repository root:

```bash
make build
./examples/simple-house/simple-house
```

Open `simple-house.bim` in a viewer that reads the dotbim schema.
