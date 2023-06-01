package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver              string `mapstructure:"DB_DRIVER"`
	DBSource              string `mapstructure:"DB_SOURCE"`
	Port                  string `mapstructure:"PORT"`
	AppID                 string `mapstructure:"APP_ID"`
	WebhookSlack          string `mapstructure:"WEBHOOK_SLACK"`
	WebhookSlackLeetCode  string `mapstructure:"WEBHOOK_SLACK_LEETCODE"`
	TagsSlackLeetCode     string `mapstructure:"TAGS_SLACK_LEETCODE"`
	ClientId              string `mapstructure:"CLIENT_ID"`
	ClientSecret          string `mapstructure:"CLIENT_SECRET"`
	RefreshToken          string `mapstructure:"REFRESH_TOKEN"`
	GrantType             string `mapstructure:"GRANT_TYPE"`
	Env                   string `mapstructure:"ENV"`
	CronNotifyRun         string `mapstructure:"CRON_NOTIFY_RUN"`
	CronNotifySummary     string `mapstructure:"CRON_NOTIFY_SUMMARY"`
	CronNotifyStatistical string `mapstructure:"CRON_NOTIFY_STATISTICAL"`
	ApiKeyUploadImage     string `mapstructure:"API_KEY_UPLOAD_IMAGE"`
	DistanceGoal          string `mapstructure:"DISTANCE_GOAL"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config *Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	env := os.Getenv("ENV")
	if env != "" {
		config.Env = env
	}
	if config.Env != "local" {
		config.Port = os.Getenv("PORT")
		config.DBSource = os.Getenv("DB_SOURCE")
		config.AppID = os.Getenv("APP_ID")
		config.WebhookSlack = os.Getenv("WEBHOOK_SLACK")
		config.WebhookSlackLeetCode = os.Getenv("WEBHOOK_SLACK_LEETCODE")
		config.TagsSlackLeetCode = os.Getenv("TAGS_SLACK_LEETCODE")
		config.ClientId = os.Getenv("CLIENT_ID")
		config.ClientSecret = os.Getenv("CLIENT_SECRET")
		config.RefreshToken = os.Getenv("REFRESH_TOKEN")
		config.GrantType = os.Getenv("GRANT_TYPE")
		config.ApiKeyUploadImage = os.Getenv("API_KEY_UPLOAD_IMAGE")
		config.DistanceGoal = os.Getenv("DISTANCE_GOAL")
	}

	return config, nil
}
