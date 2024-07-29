package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	// data содержит данные в формате gob
	data := []byte{12, 255, 129, 2, 1, 2, 255, 130, 0, 1, 12,
		0, 0, 17, 255, 130, 0, 2, 6, 72, 101, 108, 108,
		111, 44, 5, 119, 111, 114, 108, 100}

	//var buffer bytes.Buffer
	//
	//if err := gob.NewEncoder(&buffer).Encode(data); err == nil {
	//	panic(err)
	//}

	decodedData := make([]string, 0)
	var buf = bytes.NewBuffer(data)

	if err := gob.NewDecoder(buf).Decode(&decodedData); err != nil {
		panic(err)
	}

	fmt.Println(decodedData)

	// напишите код, который декодирует data в массив строк
	// 1. создайте буфер `bytes.NewBuffer(data)` для передачи в декодер
	// 2. создайте декодер `dec := gob.NewDecoder(buf)`
	// 3. определите `make([]string, 0)` для получения декодированного слайса
	// 4. декодируйте данные используя функцию `dec.Decode`

	// ...
}
