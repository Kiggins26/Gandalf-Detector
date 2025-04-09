package gandalf-discord

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Replace with actual ABI string from your contract JSON output
const contractABI = `[{"inputs":[{"internalType":"address","name":"","type":"address"}],"name":"walletToDiscordNameMapping","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

// Replace with your contract address
const contractAddress = "0xYourContractAddressHere"

func GetDiscordNameForWallet(rpcURL string, walletAddress string) (string, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return "", fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}
	defer client.Close()

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




