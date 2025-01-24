package blockchain

import (
    "time"
    "fmt"
)

// Transaction is a simple representation of a transaction.
// In real use, it will hold inputs, outputs, amounts, addresses, etc.
type Transaction struct {
    Data      string
    Timestamp time.Time
}

// Block contains basic data of a block in the chain.
type Block struct {
    Index         int
    PreviousHash  string
    Timestamp     time.Time
    Transactions  []Transaction
    Hash          string // The block's cryptographic hash
    Nonce         int64  // For proof of work or any other mechanism
}

// PrintBlock is a utility function to display block info
func (b *Block) PrintBlock() {
    fmt.Printf("Block %d:\n", b.Index)
    fmt.Printf("  PreviousHash: %s\n", b.PreviousHash)
    fmt.Printf("  Hash:         %s\n", b.Hash)
    fmt.Printf("  Nonce:        %d\n", b.Nonce)
    fmt.Printf("  Transactions: %v\n", b.Transactions)
}
