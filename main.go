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

	db, closeMongoConnection := database.OpenMongoConnection()
	defer closeMongoConnection()

	database.AddBoard(db, "TestBoard")

	fmt.Println(db)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server.InitServer()
		wg.Done()
	}()
	logger.Info("Server is now running...")

	wg.Wait()
}
