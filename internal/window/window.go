package window

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func Initialize() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	requestOpenGLContext()
}

func Create(width, height int, title string) *glfw.Window {
	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		log.Fatalln("failed to create window:", err)
	}
	centralize(window, width, height)
	return window
}

func requestOpenGLContext() {
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
}

func centralize(window *glfw.Window, width, height int) {
	// Pegar monitor principal
	monitor := glfw.GetPrimaryMonitor()
	mode := monitor.GetVideoMode()

	// Calcular posição central
	x := (mode.Width - width) / 2
	y := (mode.Height - height) / 2

	window.SetPos(x, y)
}
