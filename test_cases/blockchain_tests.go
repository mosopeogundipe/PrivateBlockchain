package test_cases

import (
	"PrivateBlockchain"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestBlockChainBasic(t *testing.T) {
	jsonBlockChain := "[{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}, {\"hash\": \"24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159\", \"timeStamp\": 1234567890, \"height\": 2, \"parentHash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"size\": 1231, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}]"
	bc := PrivateBlockchain.BlockChain{}
	bc = bc.DecodeFromJSON(jsonBlockChain)
	jsonNew := bc.EncodeToJSON()
	var realValue []PrivateBlockchain.BlockJson
	var expectedValue []PrivateBlockchain.BlockJson
	json.Unmarshal([]byte(jsonNew), &realValue)
	json.Unmarshal([]byte(jsonBlockChain), &expectedValue)
	if !reflect.DeepEqual(realValue, expectedValue) {
		fmt.Println("=========Real=========")
		fmt.Println(realValue)
		fmt.Println("=========Expcected=========")
		fmt.Println(expectedValue)
		t.Fail()
	}
}
