package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	const (
		width  = 800
		height = 800
	)

	runtime.LockOSThread()

	window := initGlfw(width, height)
	defer glfw.Terminate()

	program := initOpenGL()

	for !window.ShouldClose() {
		draw(window, program)
	}
}

func initGlfw(width int, height int) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Conway's Game of Life", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Printf("OpenGL Version: %s\n", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	var vertices []float32
	vertices = append(vertices, defTriangle([]float32{0, 0.5, 0}, []float32{-0.5, -0.5, 0}, []float32{0.5, -0.5, 0})...)

	vao := makeVao(vertices)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices))/3)

	glfw.PollEvents()
	window.SwapBuffers()
}

func defTriangle(point0 []float32, point1 []float32, point2 []float32) []float32 {
	// Build the vertex slice
	var vertices []float32
	vertices = append(vertices, point0...)
	vertices = append(vertices, point1...)
	vertices = append(vertices, point2...)

	return vertices
}

func makeVao(vertices []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}
