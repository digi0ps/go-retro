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

	database.AddBoard(db, "Board_Test_1")

	board, _ := database.FindBoard(db, "5ed15a30c7d6179a3a16345e")
	fmt.Println("Board: ", board)
	err := database.DeleteBoard(db, "5ed150534fa1b0f818b5ff57")
	fmt.Println("Delete: ", err)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server.InitServer()
		wg.Done()
	}()
	logger.Info("Server is now running...")

	wg.Wait()
}
