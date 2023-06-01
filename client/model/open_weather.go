package model

type OpenWeather struct {
	Current struct {
		Dt         float64 `json:"dt"`
		Sunrise    float64 `json:"sunrise"`
		Sunset     float64 `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   float64 `json:"pressure"`
		Humidity   float64 `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     float64 `json:"clouds"`
		Visibility float64 `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    float64 `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
}

type ParamOpenWeather struct {
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	Units   string `json:"units"`
	Appid   string `json:"appid"`
	Exclude string `json:"exclude"`
}
