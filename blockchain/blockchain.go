package blockchain

import (
	"context"
	"fmt"
	"guldo/models"
	"os"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"
	"github.com/joho/godotenv"
)

type Client struct {
	Client *rpc.Provider
}

var (
	getOddsFunctionName = "get_event_probability"
)

func NewBlockchainClient() (*Client, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	rpcURL := os.Getenv("RPC_URL")
	if rpcURL == "" {
		fmt.Println("RPC_URL environment variable is not set")
		return nil, fmt.Errorf("RPC_URL environment variable is not set")
	}

	client, err := rpc.NewProvider(rpcURL)
	if err != nil {
		fmt.Printf("Error creating blockchain client: %v", err)
		return nil, err
	}

	fmt.Println("Blockchain client created successfully")
	return &Client{Client: client}, nil
}

func (bc *Client) Call(contractAddress string, functionName string) ([]*felt.Felt, error) {
	address, err := utils.HexToFelt(contractAddress)
	if err != nil {
		fmt.Printf("Error converting address to felt: %v", err)
		return nil, err
	}

	selector := utils.GetSelectorFromNameFelt(functionName)

	call := rpc.FunctionCall{
		ContractAddress:    address,
		EntryPointSelector: selector,
		Calldata:           []*felt.Felt{},
	}

	response, err := bc.Client.Call(context.Background(), call, rpc.BlockID{Tag: "latest"})
	if err != nil {
		fmt.Printf("Error calling function: %v", err)
		return nil, err
	}

	return response, nil
}

func (bc *Client) GetEventProbability(contractAddress string) (models.OddsHistory, error) {

	response, err := bc.Call(contractAddress, getOddsFunctionName)
	if err != nil {
		return models.OddsHistory{}, err
	}

	value := []float64{}
	for _, felt := range response {
		if felt != nil {
			decimalResponse, _ := utils.FeltToBigInt(felt).Float64()
			if decimalResponse > 0 {
				res := 1 / (decimalResponse / 10000)
				value = append(value, res)
			}
		}
	}

	odd := models.OddsHistory{
		OddsYes:      value[0],
		OddsNo:       value[1],
		EventAddress: contractAddress,
	}

	return odd, nil
}
