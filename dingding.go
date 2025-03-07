package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	DINGDING_BOT_BASE_URL   = "https://oapi.dingtalk.com/robot"
	DINGDING_BOT_SEND_URL   = DINGDING_BOT_BASE_URL + "/send?access_token="
	DINGDING_BOT_UPLOAD_URL = DINGDING_BOT_BASE_URL + "/upload_media?access_token="
)

type DingDingBot struct {
	WebhookURL string
	WebhookKey string
	SignKey    string // Add this field for signature verification
}

func NewDingDingBot(webhookURL, webhookKey, signKey string) *DingDingBot {
	return &DingDingBot{
		WebhookURL: webhookURL,
		WebhookKey: webhookKey,
		SignKey:    signKey,
	}
}

// generateSignature creates a signature for DingDing API requests using HMAC-SHA256
func (bot *DingDingBot) generateSignature(timestamp int64) (string, error) {
	if bot.SignKey == "" {
		return "", nil
	}
	
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, bot.SignKey)
	
	// Create HMAC-SHA256 signature
	h := hmac.New(sha256.New, []byte(bot.SignKey))
	h.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	
	return signature, nil
}
// SendText sends a text message to the DingDing group.
func (bot *DingDingBot) SendText(content string, atMobiles []string, atUserIds []string, isAtAll bool) error {
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": content,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"atUserIds": atUserIds,
			"isAtAll":   isAtAll,
		},
	}

	return bot.sendRequest(payload)
}

// SendMarkdown sends a markdown message to the DingDing group.
func (bot *DingDingBot) SendMarkdown(title string, content string, atMobiles []string, atUserIds []string, isAtAll bool) error {
	payload := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  content,
		},
		"at": map[string]interface{}{
			"atMobiles": atMobiles,
			"atUserIds": atUserIds,
			"isAtAll":   isAtAll,
		},
	}

	return bot.sendRequest(payload)
}

// SendImage sends an image message to the DingDing group.
func (bot *DingDingBot) SendImage(base64Data string, md5 string) error {
	payload := map[string]interface{}{
		"msgtype": "image",
		"image": map[string]interface{}{
			"base64": base64Data,
			"md5":    md5,
		},
	}

	return bot.sendRequest(payload)
}

// SendNews sends a link message to the DingDing group.
func (bot *DingDingBot) SendNews(title string, text string, messageUrl string, picUrl string) error {
	payload := map[string]interface{}{
		"msgtype": "link",
		"link": map[string]interface{}{
			"title":      title,
			"text":       text,
			"messageUrl": messageUrl,
			"picUrl":     picUrl,
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

// SendTemplateCard sends an action card message to the DingDing group.
func (bot *DingDingBot) SendTemplateCard(title string, text string, singleTitle string, singleURL string, btnOrientation string) error {
	payload := map[string]interface{}{
		"msgtype": "actionCard",
		"actionCard": map[string]interface{}{
			"title":          title,
			"text":           text,
			"singleTitle":    singleTitle,
			"singleURL":      singleURL,
			"btnOrientation": btnOrientation,
		},
	}

	return bot.sendRequest(payload)
}

// UploadFile uploads a file to DingDing and returns the media ID.
func (bot *DingDingBot) UploadFile(filePath string) (string, error) {
	// Check if we're in test mode (webhook key starts with "test-")
	if len(bot.WebhookKey) >= 5 && bot.WebhookKey[:5] == "test-" {
		fmt.Printf("TEST MODE: Would upload file %s to DingDing API\n", filePath)
		return "test-media-id-12345", nil
	}

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

	requestURL := fmt.Sprintf("%s%s&type=file", bot.WebhookURL, bot.WebhookKey)
	
	// Add signature if sign key is provided
	if bot.SignKey != "" {
		timestamp := time.Now().UnixNano() / 1e6
		signature, err := bot.generateSignature(timestamp)
		if err != nil {
			return "", err
		}
		requestURL = fmt.Sprintf("%s&timestamp=%d&sign=%s", requestURL, timestamp, url.QueryEscape(signature))
	}

	resp, err := http.Post(requestURL, writer.FormDataContentType(), body)
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
		return "", fmt.Errorf("DingDing API error: %s", result["errmsg"].(string))
	}

	return result["media_id"].(string), nil
}

// sendRequest sends a request to the DingDing API with the given payload.
func (bot *DingDingBot) sendRequest(payload map[string]interface{}) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Check if we're in test mode (webhook key starts with "test-")
	if len(bot.WebhookKey) >= 5 && bot.WebhookKey[:5] == "test-" {
		fmt.Printf("TEST MODE: Would send to DingDing API: %s\n", string(jsonPayload))
		return nil
	}

	requestURL := bot.WebhookURL + bot.WebhookKey
	
	// Add signature if sign key is provided
	if bot.SignKey != "" {
		timestamp := time.Now().UnixNano() / 1e6
		signature, err := bot.generateSignature(timestamp)
		if err != nil {
			return err
		}
		requestURL = fmt.Sprintf("%s&timestamp=%d&sign=%s", requestURL, timestamp, url.QueryEscape(signature))
	}
	
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(jsonPayload))
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
		return fmt.Errorf("DingDing API error: %s", result["errmsg"].(string))
	}

	return nil
}
