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
	Normals [][]mgl64.Vec3
	DrawAsSurface bool
}

func newTerrain(scale mgl64.Vec3, heights [][]float64) *Terrain {
	t := new(Terrain)
	t.Scale = scale
	t.Heights = heights
	t.Normals = calculateNormals(scale, heights);
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

func calculateNormals(scale mgl64.Vec3, heights [][]float64) [][]mgl64.Vec3 {
	height := len(heights)
	width := len(heights[0])

	normals := make([][]mgl64.Vec3, height)
	for y:=0; y<height; y++ {
		normals[y] = make([]mgl64.Vec3, width)
		for x:=0; x<width; x++ {
			normals[y][x] = mgl64.Vec3{0.0, 1.0, 0.0}
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
		
			normals[x][y] = mgl64.Vec3{vnx, vny, vnz}
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

/*func (t *Terrain) GetTriangleUnder(pos mgl64.Vec3) Triangle {
	
}*/

func (t *Terrain) Draw() {
	h := len(t.Heights)
	w := len(t.Heights[0])
	
	gl.PushMatrix()

	var n [3]float64	
	
	gl.Translated(t.Pos[0], t.Pos[1], t.Pos[2])
	
	if t.DrawAsSurface {
		gl.Begin(gl.TRIANGLES)
		for y:=0; y<h-1; y++ {
			Y := float64(y)
			for x:=0; x<w-1; x++ {
				X := float64(x)
				gl.Vertex3d(X+0, t.Heights[y+0][x+0], Y+0)
				n = t.Normals[y+0][x+0]; gl.Normal3d(n[0], n[2], n[1])
				gl.Vertex3d(X+1, t.Heights[y+0][x+1], Y+0)
				n = t.Normals[y+0][x+1]; gl.Normal3d(n[0], n[2], n[1])
				gl.Vertex3d(X+1, t.Heights[y+1][x+1], Y+1)
				n = t.Normals[y+1][x+1]; gl.Normal3d(n[0], n[2], n[1])
				
				gl.Vertex3d(X+1, t.Heights[y+1][x+1], Y+1)
				n = t.Normals[y+1][x+1]; gl.Normal3d(n[0], n[2], n[1])
				gl.Vertex3d(X+0, t.Heights[y+1][x+0], Y+1)
				n = t.Normals[y+1][x+0]; gl.Normal3d(n[0], n[2], n[1])
				gl.Vertex3d(X+0, t.Heights[y+0][x+0], Y+0)
				n = t.Normals[y+0][x+0]; gl.Normal3d(n[0], n[2], n[1])
			}
		}
	} else {
		gl.Begin(gl.LINES)
		for y:=0; y<h-1; y++ {
			Y := float64(y)
			for x:=0; x<w-1; x++ {
				X := float64(x)
				gl.Vertex3d(X+0, t.Heights[y+0][x+0], Y+0)
				gl.Vertex3d(X+1, t.Heights[y+0][x+1], Y+0)
				gl.Vertex3d(X+1, t.Heights[y+0][x+1], Y+0)
				gl.Vertex3d(X+1, t.Heights[y+1][x+1], Y+1)
				gl.Vertex3d(X+1, t.Heights[y+1][x+1], Y+1)
				gl.Vertex3d(X+0, t.Heights[y+0][x+0], Y+0)
				gl.Vertex3d(X+1, t.Heights[y+1][x+1], Y+1)
				gl.Vertex3d(X+0, t.Heights[y+1][x+0], Y+1)
				gl.Vertex3d(X+0, t.Heights[y+1][x+0], Y+1)
				gl.Vertex3d(X+0, t.Heights[y+0][x+0], Y+0)
			}
		}
	}

	gl.End()
	
	gl.PopMatrix()
}

