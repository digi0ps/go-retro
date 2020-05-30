package main

import (
	"go-retro/config"
	"go-retro/database"
	"go-retro/database/mongodb"
	"go-retro/logger"
	"go-retro/server"
	"sync"
)

func testDatabase(db database.Service) {
	err := db.OpenConnection()
	if err != nil {
		logger.Info("Cant open database connection")
		panic(err)
	}
	defer db.CloseConnection()

	boardID, _ := db.CreateBoard("Refactor Board")
	columnID, _ := db.CreateColumn(boardID, "What went well")
	db.CreateCard(boardID, columnID, "I like the Go inteface")
	db.CreateCard(boardID, columnID, "It's helping me keep things nice")
}

func main() {
	config.LoadConfig()
	logger.Info("Config loaded...")

	mongoInstance := &mongodb.MongoDatabase{}
	testDatabase(mongoInstance)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server.InitServer()
		wg.Done()
	}()
	logger.Info("Server is now running...")

	wg.Wait()
}
