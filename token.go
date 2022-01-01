package wechats

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Token struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (c Client) GetToken() (data Token, err error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s",
		c.Corpid, c.Appsecret)

	req, err := http.Get(url)
	if err != nil {
		err = errors.New("Please check URL or Network ! # " + err.Error())
		return
	}

	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return
	}

	if data.ErrCode != 0 {
		err = errors.New("Please check corpid or corpsecret related parameter information ! # " + data.ErrMsg)
		return
	}

	return
}
