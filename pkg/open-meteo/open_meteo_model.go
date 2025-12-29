package openmeteo

type WeatherResponse struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationTimeMS     float64 `json:"generationtime_ms"`
	UTCOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`

	CurrentUnits CurrentUnits `json:"current_units"`
	Current      Current      `json:"current"`

	HourlyUnits HourlyUnits `json:"hourly_units"`
	Hourly      Hourly      `json:"hourly"`
}

type CurrentUnits struct {
	Time          string `json:"time"`
	Interval      string `json:"interval"`
	Temperature2M string `json:"temperature_2m"`
	WindSpeed10M  string `json:"wind_speed_10m"`
}

type Current struct {
	Time          string  `json:"time"`
	Interval      int     `json:"interval"`
	Temperature2M float64 `json:"temperature_2m"`
	WindSpeed10M  float64 `json:"wind_speed_10m"`
}

type HourlyUnits struct {
	Time               string `json:"time"`
	Temperature2M      string `json:"temperature_2m"`
	RelativeHumidity2M string `json:"relative_humidity_2m"`
	WindSpeed10M       string `json:"wind_speed_10m"`
}

type Hourly struct {
	Time               []string  `json:"time"`
	Temperature2M      []float64 `json:"temperature_2m"`
	RelativeHumidity2M []int     `json:"relative_humidity_2m"`
	WindSpeed10M       []float64 `json:"wind_speed_10m"`
}
