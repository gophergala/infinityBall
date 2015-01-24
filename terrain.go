package main

import (
	"fmt"
	"math"
	"github.com/go-gl/gl"
	"github.com/go-gl/mathgl/mgl64"
)

type Terrain struct {
	Scale mgl64.Vec3
	Pos mgl64.Vec3
	Heights [][]float64
	Verts [][][3]float64
	Norms [][][3]float64
	DrawAsSurface bool
}

func newTerrain(scale mgl64.Vec3, heights [][]float64) *Terrain {
	t := new(Terrain)
	t.Scale = scale
	t.Heights = heights
	t.Verts = calculateVertices(scale, heights)
	t.Norms = calculateNormals(scale, heights)
	t.DrawAsSurface = true
	return t
}

func ReadTerrain(scale mgl64.Vec3) *Terrain {
	return newTerrain(scale, readHeightmap())
}

func readHeightmap() [][]float64 {
	var n int
	fmt.Scan(&n)
	hm := make([][]float64, n)
	for y:=0; y<n; y++ {
		hm[y] = make([]float64, n)
		for x:=0; x<n; x++ {
			fmt.Scan(&hm[y][x])
		}
	}
	return hm
}

func calculateVertices(scale mgl64.Vec3, heights [][]float64) [][][3]float64 {
	h := len(heights)
	w := len(heights[0])
	verts := make([][][3]float64, h)
	for y:=0; y<h-1; y++ {
		verts[y] = make([][3]float64, w)
		for x:=0; x<w-1; x++ {
			verts[y][x] = [3]float64{float64(x), heights[y][x], float64(y)}
		}
	}
	return verts
}

func calculateNormals(scale mgl64.Vec3, heights [][]float64) [][][3]float64 {
	height := len(heights)
	width := len(heights[0])

	normals := make([][][3]float64, height)
	for y:=0; y<height; y++ {
		normals[y] = make([][3]float64, width)
		for x:=0; x<width; x++ {
			normals[y][x] = [3]float64{0.0, 0.0, 1.0}
		}
	}
	
	for y := 1; y < height-1; y++ {
		Y := float64(y)
		for x := 1; x < width-1; x++ {
			X := float64(x)
			vertex := mgl64.Vec3{X, Y, heights[x][y]}
			
			t1 := mgl64.Vec3{X-1, Y,   heights[x-1][y]}
			t2 := mgl64.Vec3{X,   Y-1, heights[x][y-1]}
			t3 := mgl64.Vec3{X+1, Y,   heights[x+1][y]}
			t4 := mgl64.Vec3{X,   Y+1, heights[x][y+1]}

			n1 := normal(vertex, t1, t2)
			n2 := normal(vertex, t2, t3)

			n3 := normal(vertex, t3, t4)
			n4 := normal(vertex, t4, t1)

			vnx := (n1[0]+n2[0]+n3[0]+n4[0])/4
			vny := (n1[1]+n2[1]+n3[1]+n4[1])/4	
			vnz := (n1[2]+n2[2]+n3[2]+n4[2])/4
		
			normals[x][y] = [3]float64{vnx, vnz, vny}
		}
	}
	return normals
}

func normal(v, t1, t2 mgl64.Vec3) mgl64.Vec3 {
	v1 := mgl64.Vec3{t1[0]-v[0], t1[1]-v[1], t1[2]-v[2]}
	v2 := mgl64.Vec3{t2[0]-v[0], t2[1]-v[1], t2[2]-v[2]}

	nx := v1[1]*v2[2] - v1[2]*v2[1]
	ny := v1[2]*v2[0] - v1[0]*v2[2]
	nz := v1[0]*v2[1] - v1[1]*v2[0]

	l := math.Sqrt(float64(nx*nx + ny*ny + nz*nz))

	return mgl64.Vec3{nx/l, ny/l, nz/l}
}

func (t *Terrain) GetTriangleUnder(pos mgl64.Vec3) Triangle {
	h := len(t.Heights)
	w := len(t.Heights[0])

	v := pos.Sub(t.Pos)
	xf,xr := math.Modf(v.X()/t.Scale.X())
	yf,yr := math.Modf(v.Z()/t.Scale.Z())
	x := int(xf)
	y := int(yf)
	
	if x<0 || y<0 || x>=w-1 || y>=h-1 {
		return Triangle{mgl64.Vec3{math.NaN(),0,0},mgl64.Vec3{0,0,0},mgl64.Vec3{0,0,0}}
	}
	
	if xr > yr {
		return Triangle{t.Verts[x][y], t.Verts[x+1][y], t.Verts[x+1][y+1]}
	} else {
		return Triangle{t.Verts[x+1][y+1], t.Verts[x][y+1], t.Verts[x][y]}
	}
}

func (t *Terrain) Draw() {
	h := len(t.Heights)
	w := len(t.Heights[0])
	
	gl.PushMatrix()

	gl.Translated(t.Pos[0], t.Pos[1], t.Pos[2])
	
	if t.DrawAsSurface {
		gl.Begin(gl.TRIANGLES)
		for y:=0; y<h-2; y++ {
			for x:=0; x<w-2; x++ {
				gl.Vertex3dv(&t.Verts[y  ][x  ])
				gl.Normal3dv(&t.Norms[y  ][x  ])
				gl.Vertex3dv(&t.Verts[y  ][x+1])
				gl.Normal3dv(&t.Norms[y  ][x+1])
				gl.Vertex3dv(&t.Verts[y+1][x+1])
				gl.Normal3dv(&t.Norms[y+1][x+1])
				gl.Vertex3dv(&t.Verts[y+1][x+1])
				gl.Normal3dv(&t.Norms[y+1][x+1])
				gl.Vertex3dv(&t.Verts[y+1][x  ])
				gl.Normal3dv(&t.Norms[y+1][x  ])
				gl.Vertex3dv(&t.Verts[y  ][x  ])
				gl.Normal3dv(&t.Norms[y  ][x  ])
			}
		}
	} else {
		gl.Begin(gl.LINES)
		for y:=0; y<h-2; y++ {
			for x:=0; x<w-2; x++ {
				gl.Vertex3dv(&t.Verts[y  ][x  ])
				gl.Vertex3dv(&t.Verts[y  ][x+1])
				gl.Vertex3dv(&t.Verts[y  ][x  ])
				gl.Vertex3dv(&t.Verts[y+1][x  ])
				gl.Vertex3dv(&t.Verts[y  ][x  ])
				gl.Vertex3dv(&t.Verts[y+1][x+1])
				gl.Vertex3dv(&t.Verts[y  ][x+1])
				gl.Vertex3dv(&t.Verts[y+1][x+1])
				gl.Vertex3dv(&t.Verts[y+1][x  ])
				gl.Vertex3dv(&t.Verts[y+1][x+1])
			}
		}
	}

	gl.End()
	
	gl.PopMatrix()
}

