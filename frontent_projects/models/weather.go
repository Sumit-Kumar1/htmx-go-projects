package models

type Weather struct {
	Latitude         *float64      `json:"latitude"`
	Longitude        *float64      `json:"longitude"`
	Elevation        *float64      `json:"elevation"`
	GenMs            *float64      `json:"generationtime_ms"`
	UTCoffsetSeconds *int          `json:"utc_offset_seconds"`
	TimeZone         *string       `json:"timezone"`
	CurrentData      *Current      `json:"current"`
	CurrentUOM       *CurrentUnits `json:"current_units"`
}

type WeatherErr struct {
	Err    *bool   `json:"error"`
	Reason *string `json:"reason"`
}

// Hourly //This struct contains temp with time fields if set any hourly fields
type Current struct {
	Time     *string  `json:"time"`
	Interval *int     `json:"interval"`
	Wind     *float64 `json:"wind_speed_10m"`
	Temp2m   *float64 `json:"temperature_2m"`
}

type CurrentUnits struct {
	Temp2m    string `json:"temperature_2m"`
	WindSpeed string `json:"wind_speed_10m"`
	Interval  string `json:"interval"`
}
