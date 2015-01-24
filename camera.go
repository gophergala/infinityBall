package main

import (
	//"math"
	"github.com/go-gl/gl"
	"github.com/go-gl/glu"
)

type Camera struct {
	X, Y, Z float64
	Roll float32
	Pitch float32
	Yaw float32
	Fov float64
}

func (cam Camera) SetupCameraLook(time float32) {
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(90, aspect, .1, 1000)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	glu.LookAt(cam.X,cam.Y,cam.Z, 0,0,0, 0,1,0)
		//fov := math.Tan(cam.Fov/360*math.Pi)
	//gl.Frustum(-fov,fov,-fov*aspect,fov*aspect, 1, 10)

	//gl.LoadIdentity()
	//gl.Translated(cam.X, cam.Y, cam.Z)
	//roll := float32(0)
	//pitch := float32(-90)
	//yaw := float32(0)
	//gl.Rotatef(roll, 0, 0, -1)
	//gl.Rotatef(pitch, 1, 0, 0)
	//gl.Translatef(0, -1, -2)
	//gl.Rotatef(yaw, 0, 1, 0)
	//gl.Translated(math.Cos(float64(time)),1,math.Sin(float64(time)))
}


