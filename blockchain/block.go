package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash     []byte //hash of this block
	Data     []byte //data to be stored
	PrevHash []byte //hash of prev block
	Nonce    int
}

// func (b *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	b.Hash = hash[:]
// }

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// badgerDb accepts only arrays of bytes or slices of bytes
// so we need to serialize and de-serialize our block data structure into bytes
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	// Package gob manages streams of gobs - binary values exchanged between an Encoder (transmitter) and a Decoder (receiver).
	// NewEncoder returns a new encoder that will transmit on the io.Writer.
	encoder := gob.NewEncoder(&res)
	// Encode transmits the data item represented by the empty interface value,
	// guaranteeing that all necessary type information has been transmitted first.
	// Passing a nil pointer to Encoder will panic, as they cannot be transmitted by gob.
	err := encoder.Encode(b)
	Handle(err)
	return res.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	Handle(err)
	return &block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
