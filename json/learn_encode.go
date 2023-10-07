package json

import (
	"encoding/json"
	"fmt"
	"os"
)

var object = []User{
	{FullName: "Fani Alfirdaus", Age: 17},
	{FullName: "John Hunt", Age: 27},
	{FullName: "Ethan Hunt", Age: 25},
	{FullName: "Joko Songo", Age: 29},
}

func EncodeFromStruct() {
	var jsonData, err = json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonString := string(jsonData)
	fmt.Println(jsonString)

	WriteToFile(jsonData)
}

func EncodeFromMap() {
	rawJson := map[string]any{
		"nama": "fani alfirdaus",
		"age":  17,
	}

	jsonData, err := json.Marshal(rawJson)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	WriteToFile(jsonData)
}

func EncodeFromAny() {
	rawJson := []any{
		map[string]any{
			"nama": "fani alfirdaus",
			"umur": 17,
		},
		map[string]any{
			"nama": "john wick",
			"umur": 25,
		},
		map[string]any{
			"nama": "wick",
			"umur": 30,
		},
	}

	jsonByte, err := json.Marshal(rawJson)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	WriteToFile(jsonByte)
}

// writing data struct to file json
func WriteToFile(data []byte) {
	path := "./data.json"
	var file *os.File

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		file, err = os.Create(path)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

	} else {
		file, err = os.OpenFile(path, os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// jika file ada isinya, maka dihapus terlebih dahulu dengan mereset size ke 0
		err = file.Truncate(0)
		if err != nil {
			fmt.Println("terjadi error saat mencoba menghapus isi file, dimana pesan error :", err.Error())
			return
		}
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("error occured on writing data where mesage is", err.Error())
		return
	}

	// simpan perubahan file secara permanent
	file.Sync()
}
