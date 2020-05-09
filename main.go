package main

import (
	"go-retro/config"
	"go-retro/logger"
	"go-retro/server"
	"sync"
)

func main() {
	logger.Info("Starting server...")

	config.LoadConfig()
	logger.Info("Config loaded...")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server.InitServer()
		wg.Done()
	}()
	logger.Info("Server is now running...")

	wg.Wait()
}
