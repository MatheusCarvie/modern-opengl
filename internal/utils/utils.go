package utils

import (
	"io"
	"log"
	"os"
)

var logFile *os.File

// InitializeLogger cria/abre app.log e redireciona o log para console + arquivo.
func InitializeLogger() {
	var err error

	logFile, err = os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Erro ao criar arquivo de log:", err)
	}

	// Escreve em console e arquivo ao mesmo tempo
	mw := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(mw)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// CloseLogger fecha o arquivo de log.
func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}
