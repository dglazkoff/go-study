package main

import (
	"reflect"
	"strconv"
	"strings"
)

// User используется для тестирования.
type User struct {
	Nick string
	Age  int `limit:"18"`
	Rate int `limit:"0,100"`
}

// Str2Int конвертирует строку в int.
func Str2Int(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

// Validate проверяет min и max для int c тегом limit.
func Validate(obj interface{}) bool {
	vobj := reflect.ValueOf(obj)
	objType := vobj.Type() // получаем описание типа

	// перебираем все поля структуры
	for i := 0; i < objType.NumField(); i++ {
		// берём значение текущего поля и проверяем, что это int
		if v, ok := vobj.Field(i).Interface().(int); ok {
			limit := objType.Field(i).Tag.Get("limit")
			limitValues := strings.Split(limit, ",")

			if v < Str2Int(limitValues[0]) {
				return false
			}

			if len(limitValues) > 1 && v > Str2Int(limitValues[1]) {
				return false
			}
		}
	}
	return true
}
