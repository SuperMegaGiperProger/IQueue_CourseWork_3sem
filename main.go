package main

import (
	"net/http"
	//	"regexp"
)

import (
	"./models"
	"./controllers/queue"
)

var routes = map[string]func(http.ResponseWriter, *http.Request) {
	"/": queue.List,
	"/create": queue.Create,
	"/push": queue.Push,
	"/show": queue.Show,
	"/save": queue.Save,
	"/remove": queue.Remove,
	"/pop": queue.Pop,
	"/list": queue.List,
}

func handleStatic() {
	fs := http.FileServer(http.Dir("static"))
  	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func initHandlers() {
	for pattern, handler := range routes {
		http.HandleFunc(pattern, handler)
	}
}

func main() {
	models.InitDB()
	handleStatic()
	initHandlers()
	http.ListenAndServe(":8080", nil)
}
