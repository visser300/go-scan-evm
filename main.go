package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"evmscanner/scripts"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	// Check if arguments were provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <contract-address1>,<contract-address2>... [block-start] [block-end]")
		fmt.Println("Example: go run main.go 0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984,0xdAC17F958D2ee523a2206206994597C13D831ec7 15000000 15100000")
		return
	}

	// Parse contract addresses (comma-separated)
	contractAddresses := strings.Split(os.Args[1], ",")

	// Parse block range
	var blockStart, blockEnd string
	if len(os.Args) > 2 {
		blockStart = os.Args[2]
	}
	if len(os.Args) > 3 {
		blockEnd = os.Args[3]
	}

	// Use the batch scanner to scan multiple contracts
	scripts.ScanMultipleContracts(contractAddresses, blockStart, blockEnd)
}
