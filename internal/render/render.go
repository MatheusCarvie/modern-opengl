package render

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v4.3-core/gl"
)

func Initialize() {
	if err := gl.Init(); err != nil {
		log.Fatalln("Failed to initialize OpenGL:", err)
	}

	fmt.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))
}

func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func SetClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func CreateVAO() *uint32 {
	var vao uint32
	gl.GenVertexArrays(1, &vao) // Cria o VAO
	gl.BindVertexArray(vao)     // Ativa o VAO
	return &vao
}

func CreateVBO(vertices []float32) *uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)              // Cria o VBO
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo) // Ativa o VBO
	gl.BufferData(                      // Envia o VBO
		gl.ARRAY_BUFFER,  // Tipo do buffer
		len(vertices)*4,  // Tamanho do buffer em bytes (4 bytes por float32)
		gl.Ptr(vertices), // Ponteiro para os dados
		gl.STATIC_DRAW,   // Modo de desenho
	)
	return &vbo
}

func CreateAttribute(index uint32, size int32, xType uint32, stride int32, offset uintptr) {
	gl.VertexAttribPointerWithOffset( // Cria um atributo
		index,  // Identificador do atributo
		size,   // Quantidade de valores (x,y)|(x,y,z)
		xType,  // Tipo dos valores
		false,  // Normalizar os valores?
		stride, // De quanto em quanto deve pular para o prox ponto x bytes
		offset, // Indica aonde começa esse atributo
	)
	gl.EnableVertexAttribArray(index) // Ativação do atribto
}
