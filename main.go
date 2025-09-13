package main

import (
	"modern-opengl/internal/render"
	"modern-opengl/internal/shader"
	"modern-opengl/internal/texture"
	"modern-opengl/internal/utils"
	"modern-opengl/internal/window"
	"runtime"

	"github.com/go-gl/gl/v4.4-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	vertices = []float32{
		/*Position*/ -0.5, -0.5 /*Color*/, 1, 0, 0 /*UV*/, 0, 0,
		/*Position*/ 0.5, -0.5 /*Color*/, 0, 1, 0 /*UV*/, 1, 0,
		/*Position*/ 0.0, 0.5 /*Color*/, 0, 0, 1 /*UV*/, 0.5, 1,
	}
)

func main() {
	// Cria files de log
	utils.InitializeLogger()
	defer utils.CloseLogger()

	// Garante que a goroutine atual fica presa em uma thread do SO
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
	render.CreateAttribute(0, 2, gl.FLOAT, 7*4, 0)   // Posição
	render.CreateAttribute(1, 3, gl.FLOAT, 7*4, 2*4) // Cor
	render.CreateAttribute(2, 2, gl.FLOAT, 7*4, 5*4) // UV

	defer gl.DeleteBuffers(1, vbo)      // Desativa o VBO
	defer gl.DeleteVertexArrays(1, vao) // Desativa o VAO

	tSource := texture.LoadTextureJPG("./assets/wood.jpg") // Carrega a textura

	vSource := shader.GetVertexSource()   // Obtem vertex source
	fSource := shader.GetFragmentSource() // Obtem o fragment source

	vShader := shader.CompileShader(gl.VERTEX_SHADER, vSource)   // Compila o vertex shader
	fShader := shader.CompileShader(gl.FRAGMENT_SHADER, fSource) // Compila o fragment shader
	shaderProgram := shader.CompileProgram(vShader, fShader)     // Compila o programa

	defer gl.DeleteShader(vShader)        // Desativa o shader
	defer gl.DeleteShader(fShader)        // Desativa o shader
	defer gl.DeleteProgram(shaderProgram) // Desativa o programa

	// Usa o programa
	gl.UseProgram(shaderProgram)

	//
	gl.Uniform1i(gl.GetUniformLocation(shaderProgram, gl.Str("texture1\x00")), 0)

	for !cw.ShouldClose() {
		render.Clear()

		gl.ActiveTexture(gl.TEXTURE0)          //
		gl.BindTexture(gl.TEXTURE_2D, tSource) //

		gl.BindVertexArray(*vao)                               // Ativa o VBO atual
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vertices)/2)) // // Desenha o VBO

		cw.SwapBuffers()
		glfw.PollEvents()
	}
}
