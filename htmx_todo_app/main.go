package main

import (
	"log"
	"net/http"

	"htmxProj/handler"
	"htmxProj/models"
)

func main() {
	t := make([]models.Tasks, 0)

	h := handler.New(t)

	http.HandleFunc("/", h.IndexPage)
	http.HandleFunc("/add", h.AddTask)
	http.HandleFunc("/delete/{id}", h.DeleteTask)
	http.HandleFunc("/done/{id}", h.Done)

	log.Fatal(http.ListenAndServe(":12344", nil))
}
