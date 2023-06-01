package weather

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/nguyenvantuan2391996/be-project/client/model"
	"github.com/sirupsen/logrus"
)

func GetWeatherInfo(url string, params *model.ParamOpenWeather) (*model.OpenWeather, error) {
	logrus.Info("Get weather information")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("lat", params.Lat)
	q.Add("lon", params.Lon)
	q.Add("units", params.Units)
	q.Add("appid", params.Appid)
	q.Add("exclude", params.Exclude)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Errorf("Close body fail %v", err)
		}
	}(res.Body)
	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		if res != nil {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				logrus.Warnf("response %v", string(bodyBytes))
				return nil, err
			}
			return nil, errors.New("get weather information failure")
		}
	}

	openWeather := &model.OpenWeather{}
	err = json.Unmarshal(body, openWeather)
	if err != nil {
		return nil, err
	}
	logrus.Info("successfully get weather information...")

	return openWeather, nil
}
