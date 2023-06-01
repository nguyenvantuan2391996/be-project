package model

import "time"

type ParamStrava struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}

type StravaToken struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	ExpiresAt    int    `json:"expires_at"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type StravaActivity struct {
	Name             string    `json:"name"`
	Distance         float64   `json:"distance"`
	MovingTime       int       `json:"moving_time"`
	Type             string    `json:"type"`
	SportType        string    `json:"sport_type"`
	ID               int64     `json:"id"`
	StartDate        time.Time `json:"start_date"`
	StartDateLocal   time.Time `json:"start_date_local"`
	Timezone         string    `json:"timezone"`
	UtcOffset        float64   `json:"utc_offset"`
	AverageSpeed     float64   `json:"average_speed"`
	MaxSpeed         float64   `json:"max_speed"`
	AverageHeartrate float64   `json:"average_heartrate"`
	MaxHeartrate     float64   `json:"max_heartrate"`
	UploadID         int64     `json:"upload_id"`
	UploadIDStr      string    `json:"upload_id_str"`
	ExternalID       string    `json:"external_id"`
}
