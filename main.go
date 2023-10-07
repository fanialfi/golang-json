package main

import (
	"fmt"

	"github.com/fanialfi/golang-json/json"
)

func main() {
	json.DecodeJsonToStruct()
	fmt.Println()
	json.DecodeJsonToMap()
	fmt.Println()
	json.DecodeJsonToAny()
	fmt.Println()
	json.ArrayJsonToArrayObj()
	fmt.Println()
	json.EncodeFromStruct()
	json.EncodeFromMap()
	json.EncodeFromAny()
}
