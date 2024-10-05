package main

import (
	"fmt"
	"guldo/blockchain"
	"guldo/db"
	"guldo/repository"
	"guldo/utils"

	"github.com/robfig/cron/v3"
)

func init() {
	utils.PrintBanner()
}

func main() {
	c := cron.New()
	c.AddFunc("@hourly", func() {
		runJob()
	})
	c.Start()

	select {} // Block forever
}

func runJob() {
	fmt.Println("Starting application...")

	database, err := connectToDatabase()
	if err != nil {
		panic(err)
	}
	defer closeDatabase(database)

	client, err := connectToBlockchain()
	if err != nil {
		panic(err)
	}

	processEvents(database, client)
}

func connectToDatabase() (*db.Database, error) {
	fmt.Println("Connecting to database...")
	database, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connected successfully")
	return database, nil
}

func closeDatabase(database *db.Database) {
	fmt.Println("Closing database connection...")
	if err := database.Close(); err != nil {
		fmt.Printf("Error closing database: %v", err)
	} else {
		fmt.Println("Database disconnected successfully")
	}
}

func connectToBlockchain() (*blockchain.Client, error) {
	fmt.Println("Connecting to blockchain...")
	client, err := blockchain.NewBlockchainClient()
	if err != nil {
		return nil, err
	}
	fmt.Println("Blockchain client connected successfully")
	return client, nil
}

func processEvents(database *db.Database, client *blockchain.Client) {
	oddsRepository := repository.NewOddsRepository(database.Conn)
	eventRepository := repository.NewEventRepository(database.Conn)

	fmt.Println("Fetching active events...")
	activeEvents, err := eventRepository.GetAllActiveEvents()
	if err != nil {
		panic(err)
	}
	fmt.Println("Active events fetched successfully")

	fmt.Println("Fetching event probability...")
	for _, contractAddress := range activeEvents {
		odds, err := client.GetEventProbability(contractAddress)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Event probability fetched successfully for contract %s: [%f : %f]\n", contractAddress, odds.OddsYes, odds.OddsNo)

		fmt.Println("Storing odds in repository...")
		err = oddsRepository.Create(odds)
		if err != nil {
			panic(err)
		}
		fmt.Println("Odds stored successfully")
	}
}
