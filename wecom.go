package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	WECOM_BOT_BASE_URL   = "https://qyapi.weixin.qq.com/cgi-bin/webhook"
	WECOM_BOT_SEND_URL   = WECOM_BOT_BASE_URL + "/send?key="
	WECOM_BOT_UPLOAD_URL = WECOM_BOT_BASE_URL + "/upload_media?key="
)

type WeComBot struct {
	WebhookURL string
	WebhookKey string
}

func NewWeComBot(webhookURL, webhookKey string) *WeComBot {
	return &WeComBot{
		WebhookURL: webhookURL,
		WebhookKey: webhookKey,
	}
}

// SendText sends a text message to the WeCom group.
func (bot *WeComBot) SendText(content string, mentionedList []string, mentionedMobileList []string) error {
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":               content,
			"mentioned_list":        mentionedList,
			"mentioned_mobile_list": mentionedMobileList,
		},
	}

	return bot.sendRequest(payload)
}

// SendMarkdown sends a markdown message to the WeCom group.
func (bot *WeComBot) SendMarkdown(content string) error {
	payload := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": content,
		},
	}

	return bot.sendRequest(payload)
}

// SendImage sends an image message to the WeCom group.
func (bot *WeComBot) SendImage(base64Data string, md5 string) error {
	payload := map[string]interface{}{
		"msgtype": "image",
		"image": map[string]interface{}{
			"base64": base64Data,
			"md5":    md5,
		},
	}

	return bot.sendRequest(payload)
}

// SendNews sends a news message to the WeCom group.
func (bot *WeComBot) SendNews(articles []NewsArticle) error {
	payload := map[string]interface{}{
		"msgtype": "news",
		"news": map[string]interface{}{
			"articles": articles,
		},
	}

	return bot.sendRequest(payload)
}

// NewsArticle represents a news article in a news message.
type NewsArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	PicURL      string `json:"picurl"`
}

// SendTemplateCard sends a template card message to the WeCom group.
func (bot *WeComBot) SendTemplateCard(cardType string, mainTitle string, mainDesc string, cardActionType int, cardActionURL string, cardActionAppID string, cardActionPagePath string) error {
	payload := map[string]interface{}{
		"msgtype": "template_card",
		"template_card": map[string]interface{}{
			"card_type": cardType,
			"main_title": map[string]interface{}{
				"title": mainTitle,
				"desc":  mainDesc,
			},
			"card_action": map[string]interface{}{
				"type":     cardActionType,
				"url":      cardActionURL,
				"appid":    cardActionAppID,
				"pagepath": cardActionPagePath,
			},
		},
	}

	return bot.sendRequest(payload)
}

// UploadFile uploads a file to WeCom and returns the media ID.
func (bot *WeComBot) UploadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("media", filePath)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return "", err
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	resp, err := http.Post(fmt.Sprintf("%s%s&type=file", bot.WebhookURL, bot.WebhookKey), writer.FormDataContentType(), body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if result["errcode"].(float64) != 0 {
		return "", fmt.Errorf("WeCom API error: %s", result["errmsg"].(string))
	}

	return result["media_id"].(string), nil
}

// sendRequest sends a request to the WeCom API with the given payload.
func (bot *WeComBot) sendRequest(payload map[string]interface{}) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	url := bot.WebhookURL + bot.WebhookKey
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	if result["errcode"].(float64) != 0 {
		return fmt.Errorf("WeCom API error: %s", result["errmsg"].(string))
	}

	return nil
}
