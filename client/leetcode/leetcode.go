package leetcode

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/nguyenvantuan2391996/be-project/client/model"
	"github.com/sirupsen/logrus"
)

func GetDailyCodingChallenge(url string, params *model.ParamDailyCodingChallenge) (*model.DailyCodingChallenge, error) {
	logrus.Info("Get daily leetcoding challenge")

	req, err := http.NewRequest("POST", url, strings.NewReader(params.Payload))
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())

	req.Header.Add("authority", "leetcode.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "en-US,en;q=0.9,vi;q=0.8,es;q=0.7")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://leetcode.com")
	req.Header.Add("referer", "https://leetcode.com/problemset/all/")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"111\", \"Not(A:Brand\";v=\"8\", \"Chromium\";v=\"111\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", model.ArrayUserAgent[rand.Intn(len(model.ArrayUserAgent)-1)])

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
			bodyBytes, errReadAll := ioutil.ReadAll(res.Body)
			if errReadAll != nil {
				logrus.Warnf("response %v", string(bodyBytes))
				return nil, errReadAll
			}
			return nil, errors.New("get daily leetcoding challenge information failure")
		}
	}

	dailyCodingChallenge := &model.DailyCodingChallenge{}
	err = json.Unmarshal(body, dailyCodingChallenge)
	if err != nil {
		return nil, err
	}
	logrus.Info("successfully get daily leetcoding challenge information...")

	return dailyCodingChallenge, nil
}
