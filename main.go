package main

import (
	"modern-opengl/internal/render"
	"modern-opengl/internal/shader"
	"modern-opengl/internal/utils"
	"modern-opengl/internal/window"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	vertices = []float32{
		/*Position*/ -0.5, -0.5 /*Color*/, 1, 0, 0,
		/*Position*/ 0.5, -0.5 /*Color*/, 0, 1, 0,
		/*Position*/ 0.0, 0.5 /*Color*/, 0, 0, 1,
	}
)

func main() {
	utils.InitializeLogger()
	defer utils.CloseLogger()

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	window.Initialize()
	defer glfw.Terminate()

	cw := window.Create(800, 600, "Modern OpenGL")
	cw.MakeContextCurrent()

	render.Initialize()
	render.SetClearColor(0, 0, 0, 1)

	vao := render.CreateVAO()
	vbo := render.CreateVBO(vertices)
	render.CreateAttribute(0, 2, gl.FLOAT, 5*4, 0)   // Posição
	render.CreateAttribute(1, 3, gl.FLOAT, 5*4, 2*4) // Cor

	defer gl.DeleteBuffers(1, vbo)      // Desativa o VBO
	defer gl.DeleteVertexArrays(1, vao) // Desativa o VAO

	vSource := shader.LoadFile("./internal/shader/vertex.glsl")   // Obtem vertex source
	fSource := shader.LoadFile("./internal/shader/fragment.glsl") // Obtem o fragment source

	vShader := shader.CompileShader(gl.VERTEX_SHADER, vSource)   // Compila o vertex shader
	fShader := shader.CompileShader(gl.FRAGMENT_SHADER, fSource) // Compila o fragment shader
	shaderProgram := shader.CompileProgram(vShader, fShader)     // Compila o programa

	defer gl.DeleteShader(vShader)        // Desativa o shader
	defer gl.DeleteShader(fShader)        // Desativa o shader
	defer gl.DeleteProgram(shaderProgram) // Desativa o programa

	gl.UseProgram(shaderProgram) // Usa o programa

	for !cw.ShouldClose() {
		render.Clear()

		gl.BindVertexArray(*vao)                               // Ativa o VBO atual
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/2)) // // Desenha o VBO

		cw.SwapBuffers()
		glfw.PollEvents()
	}
}
