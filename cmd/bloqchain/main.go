package main

import (
    "fmt"
    "log"

    "github.com/albertnieto/bloqchain/pkg/blockchain"
    "github.com/albertnieto/bloqchain/pkg/communication"
    "github.com/albertnieto/bloqchain/pkg/crypto"
    "github.com/albertnieto/bloqchain/pkg/rng"
)

func main() {
    // Initialize default cryptography module (ECDSA)
    cryptoModule := crypto.NewDefaultCrypto()

    // Initialize default RNG module (using crypto/rand for randomness)
    rngModule := rng.NewDefaultRNG()

    // Initialize secure communication module (TLS-based)
    commModule, err := communication.NewTLSChannel("server.crt", "server.key")
    if err != nil {
        log.Fatalf("Failed to initialize TLS communication: %v", err)
    }

    // Create a new blockchain
    bc, err := blockchain.NewBlockchain(cryptoModule, rngModule)
    if err != nil {
        log.Fatalf("Failed to create blockchain: %v", err)
    }

    // Start or join the network (in a real application, you'd pass addresses/peers)
    if err := commModule.StartServer(); err != nil {
        log.Fatalf("Failed to start communication server: %v", err)
    }

    // Example: Create a new block
    newBlock := bc.CreateBlock("Some transaction data")
    err = bc.AddBlock(newBlock)
    if err != nil {
        fmt.Println("Error adding block:", err)
    }

    fmt.Println("Blockchain has been initialized and a block has been added.")
    // You can expand this to handle node loops, P2P messaging, consensus, etc.
}
