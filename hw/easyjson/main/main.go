// myeasyjson/main.go
package main

import (
	"fmt"
	"github.com/mailru/easyjson"
	"gopl.io/hw/easyjson/myjson"
)

func main() {
	balance := myjson.AccountBalance{
		AccountIdHash: []byte{0x10, 0x20, 0x0A, 0x0B},
		Amounts: []myjson.CurrencyAmount{
			{Amount: 1000000, Decimals: 2, Symbol: "RUB"},
			{Amount: 2510, Decimals: 2, Symbol: "USD"},
		},
		IsBlocked: true,
	}

	// преобразуем значение переменной balance в JSON-формат
	out, err := easyjson.Marshal(balance)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
