// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example shows how to set up a minimal GLFW application.
package main

import (
	"fmt"
	"os"
	"math"
	"io/ioutil"
	"github.com/go-gl/gl"
	"github.com/go-gl/glh"
	"github.com/go-gl/glfw"
)

func exitIfError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[e] %v\n", err)
		os.Exit(1)
	}
}

func main() {
	err := glfw.Init()
	exitIfError(err)
	defer glfw.Terminate()
	
	err = glfw.OpenWindow(640, 640, 8, 8, 8, 0, 0, 0, glfw.Windowed)
	exitIfError(err)
	defer glfw.CloseWindow()

	shaderBytes, err := ioutil.ReadFile("shader.frag");
	exitIfError(err)
	shaderCode := string(shaderBytes)
	fmt.Printf("read %d bytes\n", len(shaderBytes))

	// Enable vertical sync on cards that support it.
	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle("Shady")

	glfw.SetWindowSizeCallback(onResize)
	glfw.SetMouseButtonCallback(onMouseBtn)
	glfw.SetKeyCallback(onKey)
	
	
	var time Time
	
	shader := glh.Shader{gl.FRAGMENT_SHADER, shaderCode}
	program := glh.NewProgram(shader)
	return
	
	inits()
	//program.Use()
	program.Delete()
	for glfw.Key(glfw.KeyEsc) == 0 && glfw.WindowParam(glfw.Opened) == 1 {
		
		time.Set(glfw.Time())
		handleInputs()
		render()

		// Swap front and back rendering buffers. This also implicitly calls
		// glfw.PollEvents(), so we have valid key/mouse/joystick states after
		// this. This behavior can be disabled by calling glfw.Disable with the
		// argument glfw.AutoPollEvents. You must be sure to manually call
		// PollEvents() or WaitEvents() in this case.
		glfw.SwapBuffers()
	}
}

func onResize(w, h int) {
	fmt.Printf("resized: %dx%d\n", w, h)
	gl.Viewport(0, 0, w, h)
	camera.Aspect = float64(w)/float64(h)
}

func onMouseBtn(button, state int) {
	//fmt.Printf("mouse button: %d, %d\n", button, state)
}

var keys [1024]bool

func onKey(key, state int) {
	keys[key] = state == 1;
	//fmt.Printf("key: %d, %d\n", key, state)
}



var camera Camera

func inits() {
	gl.ShadeModel(gl.SMOOTH)
	gl.ClearColor(.52, .81, .98, 0)
	gl.ClearDepth(-1)
	gl.DepthMask(true)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LEQUAL)
	//gl.DepthRangef(0,1)
	gl.Disable(gl.BLEND)
	//gl.Enable(gl.BLEND);
	//gl.BlendFunc(gl.ONE, gl.ZERO);
	gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)
}

var camYaw float64 = 4.5
var camPitch float64 = .5
const gravity float64 = -0.00982

func handleInputs() {
	if keys['W'] {camPitch += .01}
	if keys['A'] {camYaw += .01}
	if keys['S'] {camPitch -= .01}
	if keys['D'] {camYaw -= .01}
	
	camera.Pos[1] = math.Sin(camPitch)*3
	camera.Pos[0] = math.Cos(camPitch)*math.Cos(camYaw)*3
	camera.Pos[2] = math.Cos(camPitch)*math.Sin(camYaw)*3
}

func render() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	camera.SetupCameraLook()
}

