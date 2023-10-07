package json

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Age      int
	FullName string `json:"Name"`
}

var jsonString string = `{"Name":"Fani Alfirdaus","Age":17}`
var jsonData []byte = []byte(jsonString)

func DecodeJsonToStruct() {
	var data User

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("data is : %#v\n", data)
	fmt.Printf("user\t: %s\n", data.FullName)
	fmt.Printf("Age\t: %d\n", data.Age)
}

func DecodeJsonToMap() {
	data := map[string]any{}

	json.Unmarshal(jsonData, &data)
	fmt.Printf("data is : %#v\n", data)
	fmt.Printf("user\t: %s\n", data["Name"])
	fmt.Printf("Age\t: %.0f\n", data["Age"])
}

func DecodeJsonToAny() {
	var data any
	json.Unmarshal(jsonData, &data)

	dataDecode := data.(map[string]any)
	fmt.Printf("data is : %#v\n", data)
	fmt.Printf("user\t: %s\n", dataDecode["Name"])
	fmt.Printf("age\t: %.0f\n", dataDecode["Age"])
}

func ArrayJsonToArrayObj() {
	jsonString := `[
		{"Name":"Fani Alfirdaus","Age":17},
		{"Name":"John Wick","Age":25},
		{"Name":"Ethan Hunt","Age":25}
	]`

	var data []User

	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, elm := range data {
		fmt.Printf("user\t: %s\n", elm.FullName)
		fmt.Printf("Age\t: %d\n", elm.Age)
	}
}
