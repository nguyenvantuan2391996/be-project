package notify

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/nguyenvantuan2391996/be-project/client"
	"github.com/nguyenvantuan2391996/be-project/client/constants"
	"github.com/nguyenvantuan2391996/be-project/client/leetcode"
	"github.com/nguyenvantuan2391996/be-project/client/model"
	"github.com/nguyenvantuan2391996/be-project/client/strava"
	"github.com/nguyenvantuan2391996/be-project/client/upload_images"
	"github.com/nguyenvantuan2391996/be-project/client/weather"
	"github.com/nguyenvantuan2391996/be-project/config"
	internalModel "github.com/nguyenvantuan2391996/be-project/internal/domain/model"
	"github.com/nguyenvantuan2391996/be-project/internal/domain/usecase"
	"github.com/sirupsen/logrus"
)

type INotifyBotInterface interface {
	ProcessNotifyRun(ctx context.Context) error
	ProcessNotifySummary(ctx context.Context) error
	ProcessNotifyStatistical(ctx context.Context) error
	ProcessNotifyDailyLeetCodingChallenge(ctx context.Context) error
}

const (
	Location       = "Ha Noi, Viet Nam"
	Text           = "Weather information"
	Celsius        = "°C"
	FormatDateTime = "02-01-2006 15:04:05"
	FormatDate     = "02-01-2006"
	Latitude       = "21.0245"
	Longitude      = "105.8412"
	Units          = "metric"
	Exclude        = "minutely,hourly,daily,alerts"
)

type BotNotify struct {
	cfg               *config.Config
	statisticalDomain *usecase.StatisticalDomain
}

func NewBotNotify(cfg *config.Config, statisticalDomain *usecase.StatisticalDomain) *BotNotify {
	return &BotNotify{
		cfg:               cfg,
		statisticalDomain: statisticalDomain,
	}
}

func (b *BotNotify) ProcessNotifyRun(ctx context.Context) error {
	logrus.Infof(constants.BeginningTaskMessage, "ProcessNotifyRun", ctx.Value(constants.XRequestID))

	weatherInfo, err := weather.GetWeatherInfo("https://openweathermap.org/data/2.5/onecall", &model.ParamOpenWeather{
		Lat:     Latitude,
		Lon:     Longitude,
		Units:   Units,
		Appid:   b.cfg.AppID,
		Exclude: Exclude,
	})
	if err != nil {
		logrus.Errorf("get weather information from open-weather fail %v", err)
		return err
	}
	textMessage := &model.TextMessageNotifyRun{
		CurrentTime: time.Now().Format(FormatDateTime),
		Location:    Location,
		Temperature: fmt.Sprintf("%v %v feels like %v %v", weatherInfo.Current.Temp, Celsius, weatherInfo.Current.FeelsLike, Celsius),
		Weather:     weatherInfo.Current.Weather[0].Description,
		IsRunning:   "Yes",
		Note:        fmt.Sprintf("Hôm nay %v, trời không mưa nên ra ngoài thể dục, thể thao đi nhé", time.Now().Format(FormatDate)),
	}
	if weatherInfo.Current.Weather[0].Main == "Rain" {
		textMessage.IsRunning = "No"
		textMessage.Note = fmt.Sprintf("Hôm nay %v, trời mưa rồi nên ở nhà đi nhé", time.Now().Format(FormatDate))
	}
	message := &model.SlackMessage{
		Text:        Text,
		IconEmoji:   client.DefaultEmoji,
		Attachments: textMessage.ToAttachment(),
	}
	return client.SendMessageSlack(b.cfg.WebhookSlack, message)
}

func (b *BotNotify) ProcessNotifySummary(ctx context.Context) error {
	logrus.Infof(constants.BeginningTaskMessage, "ProcessNotifySummary", ctx.Value(constants.XRequestID))
	stravaActivities, err := strava.GetStravaActivityInfo(&model.ParamStrava{
		ClientId:     b.cfg.ClientId,
		ClientSecret: b.cfg.ClientSecret,
		RefreshToken: b.cfg.RefreshToken,
		GrantType:    b.cfg.GrantType,
	})
	if err != nil || len(stravaActivities) == 0 {
		return errors.New("strava activities information empty")
	}

	distanceGoal, err := strconv.ParseFloat(b.cfg.DistanceGoal, 64)
	if err != nil {
		return err
	}
	timeChart, messageActives, chartsInfo := time.Now().AddDate(0, 0, -1).Format(FormatDate), make([]*model.TextMessageNotifySummary, 0), make([]*internalModel.ChartInfo, 0)
	for _, activity := range stravaActivities {
		if timeChart != activity.StartDateLocal.Format(FormatDate) {
			continue
		}
		textMessage := &model.TextMessageNotifySummary{
			CurrentTime:      time.Now().Format(FormatDateTime),
			SportType:        activity.SportType,
			Name:             activity.Name,
			Distance:         fmt.Sprintf("%v km", math.Round((activity.Distance/float64(1000))*100)/100),
			MovingTime:       time.Duration(activity.MovingTime * 1000000000).String(),
			AverageSpeed:     fmt.Sprintf("%.2f km/h", activity.AverageSpeed*60*60/float64(1000)),
			MaxSpeed:         fmt.Sprintf("%.2f km/h", activity.MaxSpeed*60*60/float64(1000)),
			AverageHeartrate: fmt.Sprintf("%.2f bpm", activity.AverageHeartrate),
			MaxHeartrate:     fmt.Sprintf("%.2f bpm", activity.MaxHeartrate),
			Note:             fmt.Sprintf("Chúc mừng bạn đã hoàn thành mục tiêu chạy %v km ngày hôm qua %v nhé", distanceGoal, time.Now().AddDate(0, 0, -1).Format(FormatDate)),
		}
		if activity.Distance < distanceGoal*1000 {
			textMessage.Note = fmt.Sprintf("Bạn đã không hoàn thành mục tiêu chạy %v km ngày hôm qua %v rồi :sleepy: ", distanceGoal, time.Now().AddDate(0, 0, -1).Format(FormatDate))
		}
		messageActives = append(messageActives, textMessage)
		chartsInfo = append(chartsInfo, &internalModel.ChartInfo{
			Day:   timeChart,
			Value: math.Round((activity.Distance/float64(1000))*100) / 100,
		})
	}

	if len(messageActives) == 0 {
		chartsInfo = append(chartsInfo, &internalModel.ChartInfo{
			Day:   timeChart,
			Value: 0,
		})
		messageActives = append(messageActives, &model.TextMessageNotifySummary{
			CurrentTime:  time.Now().Format(FormatDateTime),
			Distance:     fmt.Sprintf("%v km", 0),
			MovingTime:   "0h0m0s",
			AverageSpeed: fmt.Sprintf("%v km/h", 0),
			MaxSpeed:     fmt.Sprintf("%v km/h", 0),
			Note:         fmt.Sprintf("Hôm qua %v, bạn không thể dục, thể thao gì à :rage:", time.Now().AddDate(0, 0, -1).Format(FormatDate)),
		})
	}

	err = b.statisticalDomain.UpsertStatistical(chartsInfo)
	if err != nil {
		return err
	}
	for _, messageActivity := range messageActives {
		message := &model.SlackMessage{
			Text:        "Activity Information",
			IconEmoji:   client.DefaultEmoji,
			Attachments: messageActivity.ToAttachment(),
		}
		err = client.SendMessageSlack(b.cfg.WebhookSlack, message)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BotNotify) ProcessNotifyStatistical(ctx context.Context) error {
	logrus.Infof(constants.BeginningTaskMessage, "ProcessNotifyStatistical", ctx.Value(constants.XRequestID))

	timeChart := strings.Join([]string{(time.Now().Month() - 1).String(), fmt.Sprintf("%v", time.Now().Year())}, "-")
	base64StringImage, sumKilometers, err := b.statisticalDomain.GetBase64StringChart(map[string]interface{}{
		"time_chart": timeChart,
	})
	if err != nil {
		return err
	}

	imageInfo, err := upload_images.UploadImage(&model.ParamUploadImage{
		Base64StringImage: base64StringImage,
		ApiKey:            b.cfg.ApiKeyUploadImage,
	})
	if err != nil {
		logrus.Errorf("get weather information from open-weather fail %v", err)
		return err
	}

	// Send message to slack
	textMessage := &model.TextMessageNotifyStatistical{
		ImageUrl: imageInfo.Image.Url,
	}

	message := &model.SlackMessage{
		Text:        strings.Join([]string{"Statistical Chart", timeChart, ". Tổng số:", fmt.Sprintf("%.2f km", sumKilometers)}, " "),
		IconEmoji:   client.DefaultEmoji,
		Attachments: textMessage.ToAttachment(),
	}
	return client.SendMessageSlack(b.cfg.WebhookSlack, message)
}

func (b *BotNotify) ProcessNotifyDailyLeetCodingChallenge(ctx context.Context) error {
	logrus.Infof(constants.BeginningTaskMessage, "ProcessNotifyDailyLeetCodingChallenge", ctx.Value(constants.XRequestID))

	dailyCodingChallenge, err := leetcode.GetDailyCodingChallenge(model.URLGraphql, &model.ParamDailyCodingChallenge{
		Payload: model.Payload,
	})
	if err != nil {
		logrus.Errorf("get daily leetcoding challenge information from leetcode.com fail %v", err)
		return err
	}

	textMsgDailyCodingChallenge := dailyCodingChallenge.ToTextMessageDailyCodingChallenge()

	message := &model.SlackMessage{
		Text:        fmt.Sprintf(model.FormatTextLeetCode, b.cfg.TagsSlackLeetCode, "This is the leetcoding challenge for today: "+dailyCodingChallenge.Data.ActiveDailyCodingChallengeQuestion.Question.Title),
		Attachments: textMsgDailyCodingChallenge.ToAttachment(),
	}

	return client.SendMessageSlack(b.cfg.WebhookSlackLeetCode, message)
}
