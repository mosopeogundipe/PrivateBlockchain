package main

import (
	"PrivateBlockchain"
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	jsonBlockChain := "[{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}, {\"hash\": \"24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159\", \"timeStamp\": 1234567890, \"height\": 2, \"parentHash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"size\": 1231, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}]"
	bc := PrivateBlockchain.BlockChain{}
	bc = bc.DecodeFromJSON(jsonBlockChain)
	fmt.Println("blockChain: ", bc)
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
		//t.Fail()
	} else {
		fmt.Println("PASSING TEST!")
	}
	//mpt := p1.MerklePatriciaTrie{}
	//mpt.Initial()
	//mpt.Insert("hello", "world")
	//mpt.Insert("charles", "ge")
	//b1 := PrivateBlockchain.Initial("", 0, mpt)
	////Initial(parent_hash string, parent_height int32, value *p1.MerklePatriciaTrie)
	//b2 := PrivateBlockchain.Initial(b1.Header.Hash, 0, mpt)
	//json1 := b1.EncodeToJSON()
	//json2 := b2.EncodeToJSON()
	////b1.DecodeFromJson(json1)
	//fmt.Println(json1)
	//fmt.Println(json2)

	//Insert(b1)
	//Insert(b2)
	//var arr []string
	//arr = append(arr, "bc")
	//arr = append(arr, "aaa")
	//str := "[" + strings.Join(arr, ",") + "]"
	//fmt.Println("Str: ", str)

	//if err != nil {
	//	fmt.Println(json1)
	//	fmt.Println(json2)
	//}
}
