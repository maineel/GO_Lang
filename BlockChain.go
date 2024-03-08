/* https://blog.logrocket.com/build-blockchain-with-go/#what-blockchain */

package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block // The first block in the blockchain
	chain        []Block
	difficulty   int // defines the minimum effort miners must undertake to mine and include a block in the blockchain.
	// This means if the mining difficulty is three, you have to generate a block hash that starts with "000" like, "0009a1bfb506…"
}

func (b Block) calculateHash() string {
	// Converted the block’s data to JSON
	data, _ := json.Marshal(b.data)

	// Concatenated the previous block’s hash, and the current block’s data, timestamp, and PoW
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)

	// Hashed the concatenated string using SHA-256
	blockHash := sha256.Sum256([]byte(blockData))

	// Returned the base 16 hash as a string
	return fmt.Sprintf("%x", blockHash)
}

// Mining is the process of adding a new block to the blockchain.
func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

func createBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0", // First block has hash 0
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) addBlock(from, to string, amount float64) {
	// collect details of the transaction
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]

	// create a new block with the transaction details
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)

	// Add the newly created block to the blockchain
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

func main() {
	// create a new blockchain instance with a mining difficulty of 2
	blockchain := createBlockchain(2)

	// record transactions on the blockchain for Alice, Bob, and John
	blockchain.addBlock("Alice", "Bob", 5)
	blockchain.addBlock("John", "Bob", 2)
	blockchain.addBlock("Neel", "Jeet", 10)
	blockchain.addBlock("Keval", "Smeet", -10)
	blockchain.addBlock("Kushal", "Zenil", 100)
	blockchain.addBlock("Adnan", "Vandit", 0)

	// check if the blockchain is valid
	fmt.Println(blockchain)
}
