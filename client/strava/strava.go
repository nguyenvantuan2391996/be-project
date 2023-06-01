package strava

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/nguyenvantuan2391996/be-project/client/model"
	"github.com/sirupsen/logrus"
)

const (
	UrlGetStravaToken    = "https://www.strava.com/oauth/token"
	UrlGetStravaActivity = "https://www.strava.com/api/v3/athlete/activities"
)

func GetStravaTokenInfo(url string, params *model.ParamStrava) (*model.StravaToken, error) {
	logrus.Info("Get strava token information")

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("client_id", params.ClientId)
	q.Add("client_secret", params.ClientSecret)
	q.Add("refresh_token", params.RefreshToken)
	q.Add("grant_type", params.GrantType)
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
			return nil, errors.New("get strava token information failure")
		}
	}

	stravaToken := &model.StravaToken{}
	err = json.Unmarshal(body, stravaToken)
	if err != nil {
		return nil, err
	}
	logrus.Info("successfully get strava token information...")

	return stravaToken, nil
}

func GetStravaActivityInfo(params *model.ParamStrava) ([]*model.StravaActivity, error) {
	logrus.Info("Get strava activity information")

	stravaToken, err := GetStravaTokenInfo(UrlGetStravaToken, params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", UrlGetStravaActivity, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("per_page", "3")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", strings.Join([]string{"Bearer ", stravaToken.AccessToken}, ""))
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			logrus.Errorf("Close body fail %v", err)
			return
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
			return nil, errors.New("get strava activity information failure")
		}
	}

	stravaActivity := make([]*model.StravaActivity, 0)
	err = json.Unmarshal(body, &stravaActivity)
	if err != nil {
		return nil, err
	}
	logrus.Info("successfully get strava activity information...")

	return stravaActivity, nil
}
