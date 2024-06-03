package main

import (
	"fmt"
	"fontent_proj/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.TemplateRun(w, "index.html")
	})

	http.HandleFunc("/pomodore", func(w http.ResponseWriter, r *http.Request) {
		handler.TemplateRun(w, "pomodore.html")
	})

	http.HandleFunc("/media", func(w http.ResponseWriter, r *http.Request) {
		handler.TemplateRun(w, "mediaPlayer.html")
	})

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		handler.TemplateRun(w, "weather.html")
	})

	http.HandleFunc("/getWeather", handler.GetWeather)

	err := http.ListenAndServe(":12344", nil)
	if err != nil {
		fmt.Printf("Error from ListenAndServe: %s", err.Error())
	}
}
