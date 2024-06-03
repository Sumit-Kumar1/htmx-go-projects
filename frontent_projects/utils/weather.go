package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"fontent_proj/models"
	"io"
	"net/http"
	"strconv"
)

func Weather(url string) (*models.Weather, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("bad Request")
	}

	var weatherResp models.Weather

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &weatherResp)
	if err != nil {
		return nil, err
	}

	emptyResp := models.Weather{}

	if weatherResp == emptyResp {
		return nil, fmt.Errorf("don't know why no resp is there")
	}

	return &weatherResp, nil
}

func WeatherToHTML(w *models.Weather) string {
	if w == nil {
		return ""
	}

	htmlStr := "<div>"

	if w.Latitude != nil {
		htmlStr += `<p>Latitude : ` + strconv.Itoa(int(*w.Latitude)) + `</p>`
	}

	if w.Longitude != nil {
		htmlStr += `<p>Longitude : ` + strconv.Itoa(int(*w.Longitude)) + `</p>`
	}

	if w.TimeZone != nil {
		htmlStr += `<p>TimeZone : ` + *w.TimeZone + `</p>`
	}

	if w.Elevation != nil {
		htmlStr += `<p>Elevation : ` + strconv.Itoa(int(*w.Elevation)) + `</p>`
	}

	if w.CurrentData != nil {
		htmlStr += `<p>Wind Speed: ` + strconv.Itoa(int(*w.CurrentData.Wind)) + " " + w.CurrentUOM.WindSpeed + "</p>"
		htmlStr += `<p>Current Temp: ` + strconv.Itoa(int(*w.CurrentData.Temp2m)) + w.CurrentUOM.Temp2m + "</p>"
	}

	htmlStr += "</div>"
	return htmlStr
}
