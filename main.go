package main

import (
	"fmt"
	"go-retro/config"
	"go-retro/database"
	"go-retro/logger"
	"go-retro/server"
	"sync"
)

func main() {
	config.LoadConfig()
	logger.Info("Config loaded...")

	pg := database.NewPostgresDB()
	pg.CreateConnection()
	defer pg.CloseConnection()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		fmt.Println(pg)
		server.InitServer()
		wg.Done()
	}()
	logger.Info("Server is now running...")

	wg.Wait()
}
