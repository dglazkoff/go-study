package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

func (a *NetAddress) Set(value string) error {
	values := strings.Split(value, `:`)
	if len(values) != 2 {
		return errors.New("Need address in a form host:port")
	}
	port, err := strconv.Atoi(values[1])
	if err != nil {
		return err
	}
	a.Host = values[0]
	a.Port = port
	return nil
}

func (a *NetAddress) String() string {
	return a.Host + ":" + strconv.Itoa(a.Port)
}

// допишите код реализации методов интерфейса
// ...

// И разберите флаг --addr=example.com:60:

func main() {
	addr := new(NetAddress)
	// если интерфейс не реализован,
	// здесь будет ошибка компиляции
	_ = flag.Value(addr)
	// проверка реализации
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}
