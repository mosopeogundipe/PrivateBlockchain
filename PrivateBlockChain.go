package PrivateBlockchain

import (
	"PrivateBlockchain/p1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/sha3"
	"strings"
	"time"
	//"unsafe"
)

type Header struct {
	Height     int32
	Timestamp  int64
	Hash       string
	ParentHash string
	Size       int32
}

type Block struct {
	Header Header
	Value  p1.MerklePatriciaTrie
}

type BlockChain struct {
	Chain  map[int32][]Block
	Length int32
}

type BlockJson struct {
	Height     int32             `json:"height"`
	Timestamp  int64             `json:"timeStamp"`
	Hash       string            `json:"hash"`
	ParentHash string            `json:"parentHash"`
	Size       int32             `json:"size"`
	MPT        map[string]string `json:"mpt"`
}

func Initial(parent_hash string, parent_height int32, value p1.MerklePatriciaTrie) Block {
	//If parent hash is empty and no block exists in blockchain, create genesis block
	if parent_hash == "" {
		return CreateGenesisBlock(value)
	}
	//Include logic to get parent's height and use to calculate block height
	height := parent_height + 1
	time_stamp := time.Now().UnixNano() / 1000000 //Current time in Unix milliseconds
	size := int32(len([]byte(value.String())))    //confirm if this approach is fine
	header := Header{Height: height, Timestamp: time_stamp, Hash: "", ParentHash: parent_hash, Size: size}
	block := Block{Header: header, Value: value}
	block.Header.Hash = block.hash_block(header, value)
	return block
}

func (b *Block) hash_block(header Header, value p1.MerklePatriciaTrie) string {
	hash_str := string(header.Height) + string(header.Timestamp) + header.ParentHash + value.GetRoot() + string(header.Size)
	hash := sha3.Sum256([]byte(hash_str))
	hash_str = hex.EncodeToString(hash[:])
	return hash_str
}

func CreateGenesisBlock(value p1.MerklePatriciaTrie) Block {
	height := int32(0)
	time_stamp := time.Now().UnixNano() / 1000000
	size := int32(len([]byte(value.String())))
	header := Header{Height: height, Timestamp: time_stamp, Hash: "", ParentHash: "0", Size: size}
	block := Block{Header: header, Value: value}
	block.Header.Hash = block.hash_block(header, value)
	return block
}

func (b *Block) EncodeToJSON() string {
	blockJsonObject := BlockJson{Height: b.Header.Height, Timestamp: b.Header.Timestamp,
		Hash: b.Header.Hash, ParentHash: b.Header.ParentHash, Size: b.Header.Size, MPT: b.Value.GetMptKeyValues()}
	blockJsonStr, _ := json.Marshal(blockJsonObject)
	return string(blockJsonStr)
}

func (b *Block) DecodeFromJson(jsonString string) Block {
	blockJsonObject := BlockJson{}
	json.Unmarshal([]byte(jsonString), &blockJsonObject)
	return convertBlockJsonToBlock(blockJsonObject)
}

func convertBlockJsonToBlock(blockJsonObject BlockJson) Block {
	blockHeader := Header{Height: blockJsonObject.Height, Timestamp: blockJsonObject.Timestamp, Hash: blockJsonObject.Hash,
		ParentHash: blockJsonObject.ParentHash, Size: blockJsonObject.Size}
	mpt := p1.MerklePatriciaTrie{}
	mpt.Initial()
	for key, value := range blockJsonObject.MPT {
		mpt.Insert(key, value)
	}
	block := Block{Header: blockHeader, Value: mpt}
	fmt.Println(block)
	return block
}

func (bc *BlockChain) EncodeToJSON() string {
	var result []string
	var blockJson string
	//fmt.Println("BC LEN: ", len(bc.Chain))
	for key := range bc.Chain {
		value := bc.Chain[key]
		//fmt.Println("LIST LEN: ", len(value))
		for index := range value {
			blockJson = value[index].EncodeToJSON()
			result = append(result, blockJson)
		}
	}
	//fmt.Println("RESULTING BC JSON: ", result)
	return "[" + strings.Join(result, ",") + "]"
}

func (bc *BlockChain) DecodeFromJSON(jsonString string) BlockChain {
	var blockJsonList []BlockJson
	json.Unmarshal([]byte(jsonString), &blockJsonList)
	//fmt.Println("BCJSON LIST OBJ: ", blockJsonList)
	var blockChain BlockChain
	for index := range blockJsonList {
		blockJson := blockJsonList[index]
		block := convertBlockJsonToBlock(blockJson)
		//fmt.Println("BLOCK OBJ:", block)
		blockChain.Insert(block)
	}
	return blockChain
}

func (bc *BlockChain) Get(height int32) []Block {
	if bc.Chain[height] == nil || len(bc.Chain[height]) == 0 {
		return nil
	} else {
		return bc.Chain[height]
	}
}

func (bc *BlockChain) Insert(block Block) {
	var blockList []Block
	if bc.Length == 0 && len(bc.Chain) == 0 {
		bc.Chain = make(map[int32][]Block)
		blockList = append(blockList, block)
		bc.Chain[bc.Length] = blockList
	} else {
		blockList = append(blockList, block)
		bc.Chain[bc.Length+1] = blockList
	}
	maxHeight := bc.FindMaxHeight()
	bc.Length = maxHeight
}

func (bc *BlockChain) FindMaxHeight() int32 {
	var maxIndex int32 = 0
	for index := range bc.Chain {
		if index > maxIndex {
			maxIndex = index
		}
	}
	return maxIndex
}
