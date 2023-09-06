package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type Counter struct {
	Value int `json:"counter"`
}

var counter = Counter{}

func (c *Counter) Increase() {
	c.Value++
}

func (c *Counter) Decrease() {
	c.Value--
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getCounterApi(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(counter)
}

func IncreaseCounterApi(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	counter.Increase()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(counter)
}

func DecreaseCounterApi(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	counter.Decrease()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(counter)
}

func getIndex(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	templ, _ := template.ParseFiles("./index.html")
	data := map[string]int{
		"Counter": counter.Value,
	}
	templ.ExecuteTemplate(w, "index.html", data)
}

func IncreaseCounter(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	tmplStr := "<div id=\"counter\">{{.Counter}}</div>"
	tmpl, _ := template.New("counter").Parse(tmplStr)
	counter.Increase()
	data := map[string]int{
		"Counter": counter.Value,
	}
	tmpl.ExecuteTemplate(w, "counter", data)
}

func DecreaseCounter(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	tmplStr := "<div id=\"counter\">{{.Counter}}</div>"
	tmpl, _ := template.New("counter").Parse(tmplStr)
	counter.Decrease()
	data := map[string]int{
		"Counter": counter.Value,
	}
	tmpl.ExecuteTemplate(w, "counter", data)
}

func main() {
	router := httprouter.New()
	router.GET("/api/", getCounterApi)
	router.POST("/api/increase", IncreaseCounterApi)
	router.POST("/api/decrease", DecreaseCounterApi)
	router.GET("/", getIndex)
	router.POST("/increase", IncreaseCounter)
	router.POST("/decrease", DecreaseCounter)

	log.Fatal(http.ListenAndServe(":8080", router))
}
