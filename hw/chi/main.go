package main

import (
	"github.com/go-chi/chi/v5"
	"io"
	"log"
	"net/http"
	"strings"
)

var cars = map[string]string{
	"id1": "Renault Logan",
	"id2": "Renault Duster",
	"id3": "BMW X6",
	"id4": "BMW M5",
	"id5": "VW Passat",
	"id6": "VW Jetta",
	"id7": "Audi A4",
	"id8": "Audi Q7",
}

// carsListFunc — вспомогательная функция для вывода всех машин.
func carsListFunc() []string {
	var list []string
	for _, c := range cars {
		list = append(list, c)
	}
	return list
}

// carFunc — вспомогательная функция для вывода определённой машины.
func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return "unknown identifier " + id
}

func carsHandle(rw http.ResponseWriter, r *http.Request) {
	carsList := carsListFunc()
	io.WriteString(rw, strings.Join(carsList, ", "))
}

func carsByBrandHandle(rw http.ResponseWriter, r *http.Request) {
	carsList := carsListFunc()
	brand := strings.ToLower(chi.URLParam(r, "brand"))
	brandCars := make([]string, 0)

	for _, car := range carsList {
		if strings.Split(strings.ToLower(car), ` `)[0] == brand {
			brandCars = append(brandCars, car)
		}
	}

	io.WriteString(rw, strings.Join(brandCars, ", "))
}

func modelHandle(rw http.ResponseWriter, r *http.Request) {
	car := strings.ToLower(chi.URLParam(r, "brand") + ` ` + chi.URLParam(r, "model"))
	for _, c := range cars {
		if strings.ToLower(c) == car {
			io.WriteString(rw, c)
			return
		}
	}
	http.Error(rw, "unknown model: "+car, http.StatusNotFound)
}

func carHandle(rw http.ResponseWriter, r *http.Request) {
	// при запросе "/car" вернётся ошибка 404,
	// поэтому не нужно проверять id на пустоту
	rw.Write([]byte(carFunc(chi.URLParam(r, "id"))))
}

func CarRouter() chi.Router {
	r := chi.NewRouter()

	r.Route("/cards", func(r chi.Router) {
		r.Get("/", carsHandle)
		r.Route("/{brand}", func(r chi.Router) {
			r.Get("/", carsByBrandHandle)
			r.Get("/{model}", modelHandle)
		})
	})

	r.Get("/car/{id}", carHandle) // GET /cars/renault/duster

	return r
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", CarRouter()))
}
