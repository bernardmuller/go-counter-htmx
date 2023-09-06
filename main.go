package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
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
	templ, err := template.ParseFiles("index.html")
	checkError(err)
	data := map[string]int{
		"Counter": counter.Value,
	}
	templ.ExecuteTemplate(w, "index.html", data)
}

func IncreaseCounter(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	tmplStr := "<div id=\"counter\">{{.Counter}}</div>"
	tmpl, err := template.New("counter").Parse(tmplStr)
	checkError(err)
	counter.Increase()
	data := map[string]int{
		"Counter": counter.Value,
	}
	tmpl.ExecuteTemplate(w, "counter", data)
}

func DecreaseCounter(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	tmplStr := "<div id=\"counter\">{{.Counter}}</div>"
	tmpl, err := template.New("counter").Parse(tmplStr)
	checkError(err)
	counter.Decrease()
	data := map[string]int{
		"Counter": counter.Value,
	}
	tmpl.ExecuteTemplate(w, "counter", data)
}

func getPort(port string) string {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}

func main() {
	router := httprouter.New()
	router.GET("/api/", getCounterApi)
	router.POST("/api/increase", IncreaseCounterApi)
	router.POST("/api/decrease", DecreaseCounterApi)
	router.GET("/", getIndex)
	router.POST("/increase", IncreaseCounter)
	router.POST("/decrease", DecreaseCounter)

	port := getPort("8080")

	log.Fatal(http.ListenAndServe("0.0.0.0"+port, router))
}
