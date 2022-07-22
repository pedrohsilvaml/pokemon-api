package main

import (
	"github.com/pedrohsilvaml/pokemon-api/internal/server"
)

func main() {
	sv := server.NewServer(&server.Config{})
	sv.Setup()

	sv.Start()
}
