package main

import (
	"fmt"

	"guldo/blockchain"
	"guldo/db"
	"guldo/repository"
	"guldo/utils"
)

func init() {
	utils.PrintBanner()
}

func main() {
	fmt.Println("Starting application...")

	fmt.Println("Connecting to database...")
	database, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully")
	oddsRepository := repository.NewOddsRepository(database.Conn)
	fmt.Println("Odds repository initialized")
	eventRepository := repository.NewEventRepository(database.Conn)
	fmt.Println("Event repository initialized")

	defer func() {
		fmt.Println("Closing database connection...")
		if err := database.Close(); err != nil {
			fmt.Printf("Error closing database: %v", err)
		} else {
			fmt.Println("Database disconnected successfully")
		}
	}()

	fmt.Println("Connecting to blockchain...")
	client, err := blockchain.NewBlockchainClient()
	if err != nil {
		panic(err)
	}
	fmt.Println("Blockchain client connected successfully")

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
