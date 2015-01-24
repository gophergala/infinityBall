// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to set up a minimal GLFW application.
package main

import (
	"fmt"
	"os"
	"math"
	"unsafe"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"github.com/go-gl/glu"
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
	for running {
		// OpenGL rendering goes here.
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



var sphere unsafe.Pointer
var terrain *Terrain
var camera Camera

func inits() {
	sphere = glu.NewQuadric()

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
	setLights(0)
	
	terrain = ReadTerrain(mgl64.Vec3{1,0.4,1})
	//terrain.DrawAsSurface = false
}

var camRotation float64

func handleInputs() {
	if keys['A'] {camRotation += .01}
	if keys['D'] {camRotation -= .01}

	camera.Y = 1
	camera.X = math.Cos(camRotation)
	camera.Z = math.Sin(camRotation)
	/*if keys['A'] {camera.X += dist}
	if keys['D'] {camera.X -= dist}
	if keys['W'] {camera.Y -= dist}
	if keys['S'] {camera.Y += dist}
	if keys['Q'] {camera.Z += dist}
	if keys['E'] {camera.Z -= dist}*/
}

func render() {
	handleInputs()
	
	time := float32(glfw.Time())
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	camera.SetupCameraLook(time)
	terrain.Draw()
	drawSphere(time)
}



func setLights(time float32) {
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

func drawSphere(time float32) {
	gl.PushMatrix()
	//gl.Rotatef(time*30,0,0,1)
	//gl.Rotatef(90,0,-1,0)
	glu.Sphere(sphere, .25, 10, 10)
	//gl.Rotatef(time*10,0,0,1)
	//gl.Translatef(.3,0,0)
	//glu.Sphere(sphere, .1, 10, 10)
	gl.PopMatrix()
	
	/*gl.PushMatrix()
	gl.Translatef(0, 0, 0)
	gl.Rectf(0,0,.3,.3)
	gl.PopMatrix()*/
}

