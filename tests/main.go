package main

import (
	"PrivateBlockchain"
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	//jsonBlockChain := "[{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}, {\"hash\": \"24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159\", \"timeStamp\": 1234567890, \"height\": 2, \"parentHash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"size\": 1231, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}]"
	jsonBlockChain := "[{\"height\":1,\"timeStamp\":1551025401,\"hash\":\"6c9aad47a370269746f172a464fa6745fb3891194da65e3ad508ccc79e9a771b\",\"parentHash\":\"genesis\",\"size\":2089,\"mpt\":{\"CS686\":\"BlockChain\",\"test1\":\"value1\",\"test2\":\"value2\",\"test3\":\"value3\",\"test4\":\"value4\"}},{\"height\":2,\"timeStamp\":1551025401,\"hash\":\"944eb943b05caba08e89a613097ac5ac7d373d863224d17b1958541088dc20e2\",\"parentHash\":\"6c9aad47a370269746f172a464fa6745fb3891194da65e3ad508ccc79e9a771b\",\"size\":2146,\"mpt\":{\"CS686\":\"BlockChain\",\"test1\":\"value1\",\"test2\":\"value2\",\"test3\":\"value3\",\"test4\":\"value4\"}},{\"height\":2,\"timeStamp\":1551025401,\"hash\":\"f8af68feadf25a635bc6e81c08f81c6740bbe1fb2514c1b4c56fe1d957c7448d\",\"parentHash\":\"6c9aad47a370269746f172a464fa6745fb3891194da65e3ad508ccc79e9a771b\",\"size\":707,\"mpt\":{\"ge\":\"Charles\"}},{\"height\":3,\"timeStamp\":1551025401,\"hash\":\"f367b7f59c651e69be7e756298aad62fb82fddbfeda26cb06bfd8adf9c8aa094\",\"parentHash\":\"f8af68feadf25a635bc6e81c08f81c6740bbe1fb2514c1b4c56fe1d957c7448d\",\"size\":707,\"mpt\":{\"ge\":\"Charles\"}},{\"height\":3,\"timeStamp\":1551025401,\"hash\":\"05ac44dd82b6cc398a5e9664add21856ae19d107d9035af5fc54c9b0ffdef336\",\"parentHash\":\"944eb943b05caba08e89a613097ac5ac7d373d863224d17b1958541088dc20e2\",\"size\":2146,\"mpt\":{\"CS686\":\"BlockChain\",\"test1\":\"value1\",\"test2\":\"value2\",\"test3\":\"value3\",\"test4\":\"value4\"}}]"
	bc := PrivateBlockchain.BlockChain{}
	bc = bc.DecodeFromJSON(jsonBlockChain)
	fmt.Println("blockChain: ", bc)
	jsonNew := bc.EncodeToJSON()
	fmt.Println("BLOCK JSON: ", jsonNew)
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
