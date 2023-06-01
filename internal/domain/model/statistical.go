package model

type Statistical struct {
	ID        int    `json:"id"`
	TimeChart string `json:"time_chart"`
	Metadata  string `json:"metadata"`
}

type ChartInfo struct {
	Day   string  `json:"day"`
	Value float64 `json:"value"`
}
