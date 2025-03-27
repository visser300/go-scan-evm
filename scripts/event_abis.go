package scripts

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Event ABI definitions
var (
	transferABI abi.ABI
	approvalABI abi.ABI
)

// Initialize ABIs
func init() {
	var err error

	// Parse Transfer event ABI
	transferABI, err = abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"}]`))
	if err != nil {
		log.Fatalf("Error parsing Transfer ABI: %v", err)
	}

	// Parse Approval event ABI
	approvalABI, err = abi.JSON(strings.NewReader(`[{"anonymous":false,"inputs":[{"indexed":true,"name":"owner","type":"address"},{"indexed":true,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"}]`))
	if err != nil {
		log.Fatalf("Error parsing Approval ABI: %v", err)
	}
}
