package main

import (
	"fmt"
	"math"
	"github.com/go-gl/mathgl/mgl64"
)

type Terrain struct {
	ScaleX float32
	ScaleY float32
	ScaleZ float32
	heights [][]int
	normals [][]mgl64.Vec3
	DrawAsSurface bool
}

func newTerrain() *Terrain {
	t := new(Terrain)
	t.ScaleX = 1
	t.ScaleY = 1
	t.ScaleZ = 1
	t.DrawAsSurface = true
	return t
}

/*
func main() {
	t := newTerrain()
	t.heights = getHeights()
	//fmt.Println(t.heights)
	calculateNormals(t)
}
*/

func calculateNormals(t *Terrain) {
	height := len(t.heights)
	width := len(t.heights[0])

	normals := make([][]mgl64.Vec3, height)
	for y:=0; y<height; y++ {
		normals[y] = make([]mgl64.Vec3, width)
		for x:=0; x<width; x++ {
			normals[y][x] = mgl64.Vec3{0.0, 1.0, 0.0}
		}
	}
	

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {

			vertex := mgl64.Vec3{float64(x), float64(y), float64(t.heights[x][y])}
			
			t1 := mgl64.Vec3{float64(x-1), float64(y), float64(t.heights[x-1][y])}
			t2 := mgl64.Vec3{float64(x), float64(y-1), float64(t.heights[x][y-1])}
			t3 := mgl64.Vec3{float64(x+1), float64(y), float64(t.heights[x+1][y])}
			t4 := mgl64.Vec3{float64(x), float64(y+1), float64(t.heights[x][y+1])}

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
	t.normals = normals;

	//fmt.Println(t.normals)
}

func normal(v, t1, t2 mgl64.Vec3) mgl64.Vec3 {
	v1 := Vec3{t1[0]-v[0], t1[1]-v[1], t1[2]-v[2]}
	v2 := Vec3{t2[0]-v[0], t2[1]-v[1], t2[2]-v[2]}

	nx := v1[1]*v2[2] - v1[2]*v2[1]
	ny := v1[2]*v2[0] - v1[0]*v2[2]
	nz := v1[0]*v2[1] - v1[1]*v2[0]

	l := math.Sqrt(float64(nx*nx + ny*ny + nz*nz))

	return mgl64.Vec3{nx/l, ny/l, nz/l}
}

func getHeights() [][]int {

	heights := make([][]int, 4)
	for y:=0; y<4; y++ {
		heights[y] = []int{1, 1, 1, 1}
	}

	return heights
	
}


