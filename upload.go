package wechats

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type UploadedMedia struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

func UploadImage(mediaPath, token string) (mediaID string, err error) {

	// body writer
	body := &bytes.Buffer{}
	bw := multipart.NewWriter(body)

	// open file
	f, err := os.Open(mediaPath)
	if err != nil {
		return
	}
	defer f.Close()

	fw, err := bw.CreateFormFile("image", mediaPath)
	if err != nil {
		return
	}

	_, err = io.Copy(fw, f)
	if err != nil {
		return
	}

	bw.Close()

	// upload image
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=image", token)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", bw.FormDataContentType())

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	var data UploadedMedia
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return
	}

	return data.MediaID, nil
}
