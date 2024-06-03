package handler

import (
	"encoding/json"
	"fmt"
	"fontent_proj/utils"
	"io"
	"net/http"
	"text/template"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	res, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while reading Request body: ", err.Error())
	}

	fmt.Println(string(res))

	url := `https://api.open-meteo.com/v1/forecast?` + string(res) + `&current=temperature_2m,wind_speed_10m`

	// get the weather from openapimetro
	resp, err := utils.Weather(url)
	if err != nil {
		fmt.Println("error while getting back resp: ", err.Error())

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	html := utils.WeatherToHTML(resp)

	w.WriteHeader(http.StatusOK)

	_, err = w.Write(json.RawMessage(html))
	if err != nil {
		fmt.Println("error while writing to response")
	}
}

func TemplateRun(w http.ResponseWriter, name string) {
	t := template.Must(template.ParseGlob(name))

	err := t.Execute(w, nil)
	if err != nil {
		fmt.Printf("Error in executing Template(%s): %s", name, err.Error())
	}
}
