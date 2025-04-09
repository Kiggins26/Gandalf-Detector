package gandalf-discord

import (
	"context"
	"fmt"
	"log"
    "os"
    "io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/joho/godotenv"
)

// Replace with actual ABI string from your contract JSON output


func GetDiscordNameForWallet(walletAddress string) (string, error) {

    err := godotenv.Load(".env")
    if err != nil{
        log.Panic("Error loading .env file: ", err)
    }
    rpcURL := os.Getenv("RPCUrl")
    contractAddress := os.Getenv("ContractAddress")
	client, err := ethclient.Dial(rpcURL)


	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

	ABIFilePath := "./abi.json"

	// Read the contents of the file
	contractABI, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic("Error reading file: ", err)
	}
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %v", err)
	}

	contractAddr := common.HexToAddress(contractAddress)
	walletAddr := common.HexToAddress(walletAddress)

	callData, err := parsedABI.Pack("walletToDiscordNameMapping", walletAddr)
	if err != nil {
		return "", fmt.Errorf("failed to pack data: %v", err)
	}

	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: callData,
	}

	ctx := context.Background()
	output, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		return "", fmt.Errorf("contract call failed: %v", err)
	}

	var discordName string
	err = parsedABI.UnpackIntoInterface(&discordName, "walletToDiscordNameMapping", output)
	if err != nil {
		return "", fmt.Errorf("failed to unpack result: %v", err)
	}

	return discordName, nil
}




