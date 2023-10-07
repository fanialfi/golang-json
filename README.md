# json data

go menyediakan package `encoding/json` untuk keperluan operasi json.

## decode json ke bentuk struct

di go data json ditulis dalam bentuk string, 
dengan menggunakan function `json.Unmarshal(data []byte, v any)` sebuah data json bisa dikonversi kedalam bentuk struct, map, maupun interface kosong.

```go
package main

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

func main(){
  var data User

  err := json.Unmarshal(jsonData, &data)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Printf("data is : %#v\n", data)
  fmt.Printf("user\t: %s\n", data.FullName)
  fmt.Printf("age\t: %d\n", data.Age)
}
```

struct `User` tersebut digunakan untuk membuat variabel baru penampung hasil decode json,
proses decoding dilakukan oleh function `json.Unmarshal()`.

Function `json.Unmarshal()` menerima data json dalam bentuk `[]byte`,
maka dari itu data json string harus di casting terlebih dahulu kebentuk `[]byte` sebelum digunakan pada function `json.Unmarshal()`. 
Lalu pada parameter kedua harus diisi dengan **pointer** dari variabel yang digunakan untuk menyimpan hasil konversi json string.

Jika dilihat lagi pada struktur struct User, properti `FullName` memiliki **tag** `json:"Name"` yang digunakan untuk mapping informasi json ke properti yang bersangkutan.
Dengan menambahkan tag json, maka properti _FullName_ pada struct, secara cerdas akan menampung data json properti _Name_.

> pada kasus decoding json string ke bentuk object struct, semua level akses pada properti struct tersebut harus bersifat public.
>> jika level akses pada properti struct tersebut tidak bersifat public, maka setelah encode data json string ke struct, pada saat pengaksesan properti pada struct akan menggunakan default value.

namun pada json string, propertinya bisa ditulis dalam bentuk huruf besar ataupun huruf kecil, go akan secara cerdas menangani tersebut.

## decode json ke bentuk `map[string]any` & `any`

tak hanya ke bentuk `struct` decoding json bisa juga ke bentuk `map[string]any` dan `any` namun untuk `any` harus di casting terlebih dahulu kebentuk `map[string]any`

```go
package main

import (
  "encoding/json"
  "fmt"
)

func main(){
  rawJson := `{"nama":"fani alfirdaus","age":17}`
  var decodeJson map[string]any

  err := json.Unmarshal([]byte(rawJson), &decodeJson)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Printf("%#v\n", decodeJson)
}
```

```go
package main

import (
  "encoding/json"
  "fmt"
)

func main(){
  rawJson := `{"nama":"fani alfirdaus","age":17}`
  var decodeJson any

  err := json.Unmarshal([]byte(rawJson), &decodeJson)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  decodeData := decodeJson.(map[])
  fmt.Printf("%#v\n", decodeData)
}
```

untuk kasus decode array json ke slice struct caranya masih sama, tinggal tipe data hasil penampung array json ke bentuk slice struct

## encode variabel struct ke json string

function `json.Marshal(v any)([]byte, error)` digunakan untuk encoding data ke json string, 
data yang akan di encode ke json string bisa berupa `struct`, `map[string]any`, atau `slice`.

Hasil konversi dari function `json.Marshal()` berupa `[]byte` dan `error` jika ada,
agar bisa ditampilkan bentuk string json-nya, harus di casting terlebih dahulu ke bentuk string.

berikut contoh encode dari struct

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
  Age int
  FullName string `json:"Name"`
}

var object = []User{
	{FullName: "Fani Alfirdaus", Age: 17},
	{FullName: "John Hunt", Age: 27},
	{FullName: "Ethan Hunt", Age: 25},
	{FullName: "Joko Songo", Age: 29},
}

func main(){
  jsonByte, err := json.Marshal(object)
  if err != nil {
    fmt.Println(err.Error())
    return
  }

  jsonString := string(jsonByte)
  fmt.Printf("%#v\n",jsonString)
}
```