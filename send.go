package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Message struct {
	Type    string
	Content interface{}
}

type Text struct {
	Content string `json:"content"`
}

func (c Client) Send(msg Message) (err error) {
	token, err := c.GetToken()
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"touser":  c.Touser,
		"toparty": c.Toparty,
		"totag":   c.Totag,
		"msgtype": msg.Type,
		"agentid": c.Agentid,
		msg.Type:  msg.Content,
	}

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token.AccessToken)
	req, err := json.Marshal(params)
	if err != nil {
		err = errors.New("Unable to marshal SMS request parameters ! # " + err.Error())
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		err = errors.New("Send text message request exception ! # " + err.Error())
		return
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return
	}

	if data["errmsg"] != "ok" {
		return errors.New(fmt.Sprintf("Message push failed ! # %v", data))
	}

	return
}

func (c Client) SendText(content string) (err error) {

	msg := Message{
		Type: "text",
		Content: Text{
			Content: content,
		},
	}

	return c.Send(msg)

}

func (c Client) SendTextOld(context string) (err error) {

	token, err := c.GetToken()
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"touser":  c.Touser,
		"toparty": c.Toparty,
		"totag":   c.Totag,
		"msgtype": "text",
		"agentid": c.Agentid,
		"text": map[string]interface{}{
			"content": context,
		},
	}

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token.AccessToken)
	req, err := json.Marshal(params)
	if err != nil {
		err = errors.New("Unable to marshal SMS request parameters ! # " + err.Error())
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		err = errors.New("Send text message request exception ! # " + err.Error())
		return
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return
	}

	if data["errmsg"] != "ok" {
		return errors.New(fmt.Sprintf("Message push failed ! # %v", data))
	}

	return
}

func (c Client) SendImage(mediaID string) (err error) {

	token, err := c.GetToken()
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"touser":  c.Touser,
		"toparty": c.Toparty,
		"totag":   c.Totag,
		"msgtype": "image",
		"agentid": c.Agentid,
		"image": map[string]interface{}{
			"media_id": mediaID,
		},
	}

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token.AccessToken)
	req, err := json.Marshal(params)
	if err != nil {
		err = errors.New("Unable to marshal SMS request parameters ! # " + err.Error())
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		err = errors.New("Send text message request exception ! # " + err.Error())
		return
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return
	}

	if data["errmsg"] != "ok" {
		return errors.New(fmt.Sprintf("Message push failed ! # %v", data))
	}

	return
}

func (c Client) SendFile(mediaID string) (err error) {

	token, err := c.GetToken()
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"touser":  c.Touser,
		"toparty": c.Toparty,
		"totag":   c.Totag,
		"msgtype": "file",
		"agentid": c.Agentid,
		"file": map[string]interface{}{
			"media_id": mediaID,
		},
	}

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token.AccessToken)
	req, err := json.Marshal(params)
	if err != nil {
		err = errors.New("Unable to marshal SMS request parameters ! # " + err.Error())
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		err = errors.New("Send text message request exception ! # " + err.Error())
		return
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return
	}

	if data["errmsg"] != "ok" {
		return errors.New(fmt.Sprintf("Message push failed ! # %v", data))
	}

	return
}

type Card struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Btntxt      string `json:"btntxt"`
}

func (c Client) SendCard(title, description, url string) (err error) {

	msg := Message{
		Type: "textcard",
		Content: Card{
			Title:       title,
			Description: description,
			Url:         url,
		},
	}

	return c.Send(msg)

}

func (c Client) SendCardOld(textcard Card) (err error) {

	token, err := c.GetToken()
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"touser":   c.Touser,
		"toparty":  c.Toparty,
		"totag":    c.Totag,
		"msgtype":  "textcard",
		"agentid":  c.Agentid,
		"textcard": textcard,
	}

	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s", token.AccessToken)
	req, err := json.Marshal(params)
	if err != nil {
		err = errors.New("Unable to marshal SMS request parameters ! # " + err.Error())
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	if err != nil {
		err = errors.New("Send text message request exception ! # " + err.Error())
		return
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return
	}

	if data["errmsg"] != "ok" {
		return errors.New(fmt.Sprintf("Message push failed ! # %v", data))
	}

	return
}
