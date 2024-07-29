package main

import (
	"fmt"
	"reflect"
)

//import (
//	"fmt"
//	"github.com/caarlos0/env/v6"
//	"log"
//)
//
//type Config struct {
//	User string `env:"USER"`
//}
//
//func main() {
//	var cfg Config
//	err := env.Parse(&cfg)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("Current user is %s\n", cfg.User)
//}

// Собственный тип для ошибок. Аналогичен стандартному error
//
//	type Error interface {
//		Error() string
//	}
//
// // MyError — структура, реализующая нашу ошибку
//
//	type MyError struct {
//		// ...
//	}
//
// // Error — метод для удовлетворения интерфейсу Error
//
//	func (e *MyError) Error() string {
//		return "..."
//	}
//
// // переменная типа ошибки — указатель на пустую структуру
// var ErrFriday13 = &MyError{}
//
//	func CheckTodayIsOkay() Error {
//		var err *MyError
//
//		t := time.Now()
//		if t.Day() == 28 {
//			err = ErrFriday13
//		}
//
//		return err
//	}
//
//	func main() {
//		err := CheckTodayIsOkay()
//		fmt.Println(err)
//		// проверяем err на nil — и внезапно всегда не nil
//		if err != nil {
//			fmt.Println("error is not nil")
//			return
//		}
//
//		fmt.Println("error is nil")
//	}
type MyType struct{}

func NaiveIsNil(obj interface{}) bool {
	return obj == nil
}

func main() {
	var t *MyType
	fmt.Println(t == nil)
	fmt.Printf("Проверка типа (%v) на nil: %v\n", reflect.TypeOf(t), NaiveIsNil(t)) // TypeOf возвращает тип переданного объекта.
}
