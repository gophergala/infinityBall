// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to set up a minimal GLFW application.
package main

import (
	"fmt"
	"os"
	"math"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"github.com/go-gl/mathgl/mgl64"
)

func main() {
	var err error
	if err = glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return
	}

	// Ensure glfw is cleanly terminated on exit.
	defer glfw.Terminate()

	if err = glfw.OpenWindow(640, 640, 8, 8, 8, 0, 0, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		return
	}

	// Ensure window is cleanly closed on exit.
	defer glfw.CloseWindow()

	// Enable vertical sync on cards that support it.
	glfw.SetSwapInterval(1)

	// Set window title
	glfw.SetWindowTitle("Simple GLFW window")

	// Hook some events to demonstrate use of callbacks.
	// These are not necessary if you don't need them.
	glfw.SetWindowSizeCallback(onResize)
	glfw.SetWindowCloseCallback(onClose)
	glfw.SetMouseButtonCallback(onMouseBtn)
	glfw.SetMouseWheelCallback(onMouseWheel)
	glfw.SetKeyCallback(onKey)
	glfw.SetCharCallback(onChar)
	
	
	inits()
	// Start loop
	running := true
	var time Time
	for running {
		time.Set(glfw.Time())
		handleInputs()
		physics(time)
		render()

		// Swap front and back rendering buffers. This also implicitly calls
		// glfw.PollEvents(), so we have valid key/mouse/joystick states after
		// this. This behavior can be disabled by calling glfw.Disable with the
		// argument glfw.AutoPollEvents. You must be sure to manually call
		// PollEvents() or WaitEvents() in this case.
		glfw.SwapBuffers()

		// Break out of loop when Escape key is pressed, or window is closed.
		running = glfw.Key(glfw.KeyEsc) == 0 && glfw.WindowParam(glfw.Opened) == 1
	}
}

var aspect float64

func onResize(w, h int) {
	fmt.Printf("resized: %dx%d\n", w, h)
	gl.Viewport(0, 0, w, h)
	aspect = float64(w)/float64(h)
}

func onClose() int {
	fmt.Println("closed")
	return 1 // return 0 to keep window open.
}

func onMouseBtn(button, state int) {
	fmt.Printf("mouse button: %d, %d\n", button, state)
}

func onMouseWheel(delta int) {
	fmt.Printf("mouse wheel: %d\n", delta)
}

var keys [1024]bool

func onKey(key, state int) {
	keys[key] = state == 1;
	fmt.Printf("key: %d, %d\n", key, state)
}

func onChar(key, state int) {
	fmt.Printf("char: %d, %d\n", key, state)
}



var terrain *Terrain
var camera Camera
var ball *Ball

func inits() {
	ball = NewBall()
	terrain = ReadTerrain(mgl64.Vec3{1,0.4,1})
	//terrain.DrawAsSurface = false
	
	gl.ShadeModel (gl.SMOOTH)
	gl.ClearColor (0.0, 0.0, 0.0, 0.0)
	gl.ClearDepth(-1.0)
	gl.DepthMask(true)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	//gl.DepthRangef(0,1)
	gl.Disable(gl.BLEND)
	gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)

	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)
	setLights()
}

var camYaw float64
var camPitch float64 = .5

func handleInputs() {
	if keys['A'] {camYaw += .01}
	if keys['D'] {camYaw -= .01}
	if keys['W'] {camPitch += .01}	
	if keys['S'] {camPitch -= .01}	

	camera.Y = math.Sin(camPitch)*3
	camera.X = math.Cos(camPitch)*math.Cos(camYaw)*3
	camera.Z = math.Cos(camPitch)*math.Sin(camYaw)*3
	/*if keys['A'] {camera.X += dist}
	if keys['D'] {camera.X -= dist}
	if keys['W'] {camera.Y -= dist}
	if keys['S'] {camera.Y += dist}
	if keys['Q'] {camera.Z += dist}
	if keys['E'] {camera.Z -= dist}*/
}

func physics(time Time) {
	
}

func render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	camera.SetupCameraLook()
	terrain.Draw()
	ball.Draw()
}



func setLights() {
	whiteSpecularLight := []float32{ 1,1,1 }
	blackAmbientLight := []float32{ 0,0,0 }
	whiteDiffuseLight := []float32{ 1,1,1 }

	mat_specular := []float32{1.0, 1.0, 1.0}
	mat_shininess := []float32{ 50.0 }
	light_position := []float32{ 0.0, 10.0, 1 }
	
	gl.Lightfv(gl.LIGHT0, gl.SPECULAR, whiteSpecularLight);
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, blackAmbientLight);
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, whiteDiffuseLight);
	gl.Lightfv(gl.LIGHT0, gl.POSITION, light_position)

	gl.Materialfv(gl.FRONT, gl.SPECULAR, mat_specular)
	gl.Materialfv(gl.FRONT, gl.SHININESS, mat_shininess)
}

