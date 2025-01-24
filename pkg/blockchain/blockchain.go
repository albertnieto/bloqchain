package blockchain

import (
    "crypto/sha256"
    "encoding/hex"
    "errors"
    "time"

    "github.com/albertnieto/bloqchain/pkg/crypto"
    "github.com/albertnieto/bloqchain/pkg/rng"
)

// Blockchain holds the chain of blocks and interfaces for cryptography, RNG
type Blockchain struct {
    Chain        []Block
    CryptoModule crypto.Crypto
    RNGModule    rng.RNG
}

// NewBlockchain initializes the blockchain with a genesis block
func NewBlockchain(cryptoModule crypto.Crypto, rngModule rng.RNG) (*Blockchain, error) {
    bc := &Blockchain{
        CryptoModule: cryptoModule,
        RNGModule:    rngModule,
    }

    // Create a genesis block
    genesisBlock := Block{
        Index:        0,
        PreviousHash: "0",
        Timestamp:    time.Now(),
        Transactions: []Transaction{{Data: "Genesis Block", Timestamp: time.Now()}},
        Nonce:        0,
    }
    genesisBlock.Hash = bc.calculateHash(genesisBlock)
    bc.Chain = append(bc.Chain, genesisBlock)

    return bc, nil
}

// CreateBlock creates a new block with random Nonce
func (bc *Blockchain) CreateBlock(data string) Block {
    lastBlock := bc.Chain[len(bc.Chain)-1]
    newBlock := Block{
        Index:        len(bc.Chain),
        PreviousHash: lastBlock.Hash,
        Timestamp:    time.Now(),
        Transactions: []Transaction{{Data: data, Timestamp: time.Now()}},
        // Nonce: Could use bc.RNGModule to generate random number
        Nonce: int64(bc.RNGModule.Intn(1_000_000_000)),
    }
    newBlock.Hash = bc.calculateHash(newBlock)
    return newBlock
}

// AddBlock adds a new block to the chain after validating it
func (bc *Blockchain) AddBlock(newBlock Block) error {
    lastBlock := bc.Chain[len(bc.Chain)-1]

    if newBlock.PreviousHash != lastBlock.Hash {
        return errors.New("invalid block: previous hash does not match")
    }

    if newBlock.Hash != bc.calculateHash(newBlock) {
        return errors.New("invalid block: hash mismatch")
    }

    bc.Chain = append(bc.Chain, newBlock)
    return nil
}

// calculateHash is a simple SHA-256 hash function for demonstration
func (bc *Blockchain) calculateHash(block Block) string {
    // In practice, you'd include the full set of transaction data, Nonce, etc.
    record := string(block.Index) + block.PreviousHash + block.Timestamp.String() + block.Transactions[0].Data + string(block.Nonce)
    hash := sha256.Sum256([]byte(record))
    return hex.EncodeToString(hash[:])
}
