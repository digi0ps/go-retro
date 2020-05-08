package main

import (
	"go-retro/logger"
	"go-retro/server"
)

func main() {
	logger.Info("Starting server...")

	server.InitServer()
}
