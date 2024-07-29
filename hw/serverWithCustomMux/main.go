package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Subj struct {
	Product string `json:"name"`
	Price   int    `json:"price"`
}

func JSONHandler(w http.ResponseWriter, req *http.Request) {
	// собираем данные
	subj := Subj{"Milk", 50}
	// кодируем в JSON
	resp, err := json.Marshal(subj)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// устанавливаем заголовок Content-Type
	// для передачи клиенту информации, кодированной в JSON
	w.Header().Set("content-type", "application/json")
	// устанавливаем код 200
	w.WriteHeader(http.StatusOK)
	// пишем тело ответа
	w.Write(resp)
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	body := fmt.Sprintf("Method: %s\r\n", req.Method)
	body += "Header ===============\r\n"
	for k, v := range req.Header {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	body += "Query parameters ===============\r\n"

	req.ParseForm()
	for k, v := range req.Form {
		body += fmt.Sprintf("%s: %v\r\n", k, v)
	}
	res.Write([]byte(body))
}

func mainFileHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./main.go")
}

func parentDirectoryHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "../")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, mainFileHandler)

	fs := http.FileServer(http.Dir("../"))
	mux.Handle(`/golang/`, http.StripPrefix(`/golang/`, fs))

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
