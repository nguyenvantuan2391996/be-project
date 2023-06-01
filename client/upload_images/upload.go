package upload_images

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/nguyenvantuan2391996/be-project/client/model"
	"github.com/sirupsen/logrus"
)

const (
	UrlHost = "https://freeimage.host/api/1/upload"
)

func UploadImage(params *model.ParamUploadImage) (*model.ImageInfo, error) {
	logrus.Info("Get link upload image")

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("source", params.Base64StringImage)
	_ = writer.WriteField("key", params.ApiKey)
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", UrlHost, payload)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		if res != nil {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				logrus.Warnf("response %v", string(bodyBytes))
				return nil, err
			}
			return nil, errors.New("link upload image failure")
		}
	}

	imgInfo := &model.ImageInfo{}
	err = json.Unmarshal(body, imgInfo)
	if err != nil {
		return nil, err
	}
	logrus.Info("successfully link upload image...")

	return imgInfo, nil
}
