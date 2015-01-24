package main

import (
	"fmt"
	"github.com/go-gl/gl"
)

type Terrain struct {
	ScaleX float32
	ScaleY float32
	ScaleZ float32
	heights [][]int
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

func ReadTerrain() *Terrain {
	t := newTerrain()
	t.heights = readHeightmap()
	return t
}

func readHeightmap() [][]int {
	var n int
	fmt.Scan(&n)
	hm := make([][]int, n)
	for y:=0; y<n; y++ {
		hm[y] = make([]int, n)
		for x:=0; x<n; x++ {
			fmt.Scan(&hm[y][x])
		}
	}
	return hm
}


func (t *Terrain) Draw() {
	h := len(t.heights)
	w := len(t.heights[0])
	
	gl.PushMatrix()
	gl.Translatef(-float32(w)/2, 0, -float32(h)/2.0)
	//gl.Scalef(t.ScaleX, t.ScaleY, t.ScaleZ)
	
	if t.DrawAsSurface {
		gl.Begin(gl.QUADS)
		for y:=0; y<h-1; y++ {
			for x:=0; x<w-1; x++ {
				gl.Vertex3i(x+0, t.heights[y+0][x+0], y+0)
				gl.Normal3i(0,1,0)
				gl.Vertex3i(x+0, t.heights[y+1][x+0], y+1)
				gl.Normal3i(0,1,0)
				gl.Vertex3i(x+1, t.heights[y+1][x+1], y+1)
				gl.Normal3i(0,1,0)
				gl.Vertex3i(x+1, t.heights[y+0][x+1], y+0)
				gl.Normal3i(0,1,0)
			}
		}
	} else {
		gl.Begin(gl.LINES)
		for y:=0; y<h-1; y++ {
			for x:=0; x<w-1; x++ {
				gl.Vertex3i(x+0, y+0, t.heights[y+0][x+0])
				gl.Vertex3i(x+1, y+0, t.heights[y+0][x+1])
				gl.Vertex3i(x+1, y+0, t.heights[y+0][x+1])
				gl.Vertex3i(x+1, y+1, t.heights[y+1][x+1])
				gl.Vertex3i(x+1, y+1, t.heights[y+1][x+1])
				gl.Vertex3i(x+0, y+0, t.heights[y+0][x+0])
				gl.Vertex3i(x+1, y+1, t.heights[y+1][x+1])
				gl.Vertex3i(x+0, y+1, t.heights[y+1][x+0])
				gl.Vertex3i(x+0, y+1, t.heights[y+1][x+0])
				gl.Vertex3i(x+0, y+0, t.heights[y+0][x+0])
			}
		}
	}

	gl.End()
	
	gl.PopMatrix()
}

