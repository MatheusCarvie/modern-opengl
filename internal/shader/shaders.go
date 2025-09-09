package shader

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gl/gl/v4.5-core/gl"
)

// LoadFile carrega o código fonte de um shader a partir de um arquivo.
// Se falhar, encerra o programa.
func LoadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Erro ao ler shader %s: %v", path, err)
	}
	return string(data)
}

// CompileShader compila um shader de um tipo específico (vertex/fragment/etc.).
// Se falhar, encerra o programa.
func CompileShader(xType uint32, source string) uint32 {
	cs := gl.CreateShader(xType) // Cria o objeto Shader

	// Converte a string em C string
	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(cs, 1, csources, nil)
	free() // libera memória temporária

	gl.CompileShader(cs) // Compila o shader

	if err := checkShaderCompile(cs); err != nil {
		log.Fatalf("Falha ao compilar shader (%d): %v", xType, err)
	}
	return cs
}

// CompileProgram cria um programa a partir de um vertex shader e um fragment shader.
// Se falhar, encerra o programa.
func CompileProgram(vertexShader, fragmentShader uint32) uint32 {
	program := gl.CreateProgram()            // Cria o programa
	gl.AttachShader(program, vertexShader)   // Anexa o vertex shader
	gl.AttachShader(program, fragmentShader) // Anexa o fragment shader
	gl.LinkProgram(program)                  // Linka os shaders

	if err := checkProgramLink(program); err != nil {
		log.Fatalf("Falha ao linkar programa: %v", err)
	}
	return program
}

// checkShaderCompile verifica se o shader compilou com sucesso
func checkShaderCompile(shader uint32) error {
	var success int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logMsg := make([]byte, logLength+1)
		gl.GetShaderInfoLog(shader, logLength, nil, &logMsg[0])

		return fmt.Errorf("erro de compilação do shader: %s", string(logMsg))
	}
	return nil
}

// checkProgramLink verifica se o programa foi linkado com sucesso
func checkProgramLink(program uint32) error {
	var success int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		logMsg := make([]byte, logLength+1)
		gl.GetProgramInfoLog(program, logLength, nil, &logMsg[0])

		return fmt.Errorf("erro ao linkar programa: %s", string(logMsg))
	}
	return nil
}
