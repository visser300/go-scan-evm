# EVM Contract Event Scanner

A Go-based project for scanning and analyzing EVM blockchain events, particularly token transfers.

## Setup

1. Clone this repository
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Copy `.env.example` to `.env` and add your Ethereum RPC URL:
   ```
   ETH_RPC_URL=https://mainnet.infura.io/v3/YOUR_INFURA_KEY
   ```

## Available Scripts

### Scan Transfer Events

Scans for ERC20 transfer events for a specific token contract:
