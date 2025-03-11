// Package main provides functionality for interacting with DingDing Bot API
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

// DingDing API endpoints
const (
	// DINGDING_BOT_BASE_URL is the base URL for DingDing Bot API
	DINGDING_BOT_BASE_URL = "https://oapi.dingtalk.com/robot"
	
	// DINGDING_BOT_SEND_URL is the endpoint for sending messages
	DINGDING_BOT_SEND_URL = DINGDING_BOT_BASE_URL + "/send?access_token="
	
	// DINGDING_BOT_UPLOAD_URL is the endpoint for uploading media files
	DINGDING_BOT_UPLOAD_URL = DINGDING_BOT_BASE_URL + "/upload_media?access_token="
)

// DingDingBot represents a DingDing Bot instance with configuration for API access
type DingDingBot struct {
	// WebhookURL is the base URL for the DingDing Bot API
	WebhookURL string
	
	// WebhookKey is the access token for the DingDing Bot
	WebhookKey string
	
	// SignKey is the secret key used for signature verification
	// This is optional but recommended for enhanced security
	SignKey string
}

// NewDingDingBot creates a new DingDingBot instance with the provided configuration
// Parameters:
//   - webhookURL: The base URL for the DingDing Bot API
//   - webhookKey: The access token for the DingDing Bot
//   - signKey: The secret key for signature verification (optional)
// Returns:
//   - A pointer to a new DingDingBot instance
func NewDingDingBot(webhookURL, webhookKey, signKey string) *DingDingBot {
	return &DingDingBot{
		WebhookURL: webhookURL,
		WebhookKey: webhookKey,
		SignKey:    signKey,
	}
}

// generateSignature creates a signature for DingDing API requests using HMAC-SHA256
// The signature is used to verify that requests are coming from authorized sources
// Parameters:
//   - timestamp: The current timestamp in milliseconds
// Returns:
//   - The Base64-encoded HMAC-SHA256 signature
//   - An error if signature generation fails
// Note:
//   - Returns an empty string and nil error if SignKey is not provided
func (bot *DingDingBot) generateSignature(timestamp int64) (string, error) {
	// If no sign key is provided, return empty signature
	if bot.SignKey == "" {
		return "", nil
	}
	
	// Format the string to sign: timestamp + newline + secret
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, bot.SignKey)
	
	// Create HMAC-SHA256 signature
	h := hmac.New(sha256.New, []byte(bot.SignKey))
	if _, err := h.Write([]byte(stringToSign)); err != nil {
		return "", fmt.Errorf("failed to create signature: %v", err)
	}
	
	// Encode the signature as Base64
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	
	return signature, nil
}
// SendText sends a text message to the DingDing group.
// Parameters:
//   - content: The text content of the message
//   - atMobiles: Array of mobile numbers to @mention
//   - atUserIds: Array of user IDs to @mention
//   - isAtAll: Whether to @mention all members in the group
// Returns:
//   - An error if the request fails, nil otherwise
func (bot *DingDingBot) SendText(content string, atMobiles []string, atUserIds []string, isAtAll bool) error {
	if content == "" {
		return fmt.Errorf("content cannot be empty")
	}
	
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
// Parameters:
//   - title: The title of the markdown message
//   - content: The markdown content of the message
//   - atMobiles: Array of mobile numbers to @mention
//   - atUserIds: Array of user IDs to @mention
//   - isAtAll: Whether to @mention all members in the group
// Returns:
//   - An error if the request fails, nil otherwise
func (bot *DingDingBot) SendMarkdown(title string, content string, atMobiles []string, atUserIds []string, isAtAll bool) error {
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if content == "" {
		return fmt.Errorf("content cannot be empty")
	}
	
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
// Parameters:
//   - base64Data: The Base64-encoded image data
//   - md5: The MD5 hash of the image
// Returns:
//   - An error if the request fails, nil otherwise
func (bot *DingDingBot) SendImage(base64Data string, md5 string) error {
	if base64Data == "" {
		return fmt.Errorf("base64Data cannot be empty")
	}
	if md5 == "" {
		return fmt.Errorf("md5 cannot be empty")
	}
	
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
// Parameters:
//   - title: The title of the news message
//   - text: The text description of the news
//   - messageUrl: The URL to open when clicking on the news
//   - picUrl: The URL of the image to display in the news
// Returns:
//   - An error if the request fails, nil otherwise
func (bot *DingDingBot) SendNews(title string, text string, messageUrl string, picUrl string) error {
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if messageUrl == "" {
		return fmt.Errorf("messageUrl cannot be empty")
	}
	
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
// This struct is used to define the structure of a news article
// when sending news messages to DingDing.
type NewsArticle struct {
	// Title is the title of the news article
	Title string `json:"title"`
	
	// Description is the text description of the news article
	Description string `json:"description"`
	
	// URL is the link that will be opened when clicking on the news article
	URL string `json:"url"`
	
	// PicURL is the URL of the image to display in the news article
	PicURL string `json:"picurl"`
}

// SendTemplateCard sends an action card message to the DingDing group.
// Parameters:
//   - title: The title of the template card
//   - text: The text content of the template card
//   - singleTitle: The title of the single button
//   - singleURL: The URL to open when clicking the button
//   - btnOrientation: The orientation of buttons ("0" for vertical, "1" for horizontal)
// Returns:
//   - An error if the request fails, nil otherwise
func (bot *DingDingBot) SendTemplateCard(title string, text string, singleTitle string, singleURL string, btnOrientation string) error {
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if text == "" {
		return fmt.Errorf("text cannot be empty")
	}
	if singleTitle == "" {
		return fmt.Errorf("singleTitle cannot be empty")
	}
	if singleURL == "" {
		return fmt.Errorf("singleURL cannot be empty")
	}
	
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
// Parameters:
//   - filePath: The path to the file to upload
// Returns:
//   - The media ID of the uploaded file, which can be used in other API calls
//   - An error if the upload fails, nil otherwise
func (bot *DingDingBot) UploadFile(filePath string) (string, error) {
	if filePath == "" {
		return "", fmt.Errorf("filePath cannot be empty")
	}
	
	// Check if we're in test mode (webhook key starts with "test-")
	if len(bot.WebhookKey) >= 5 && bot.WebhookKey[:5] == "test-" {
		fmt.Printf("TEST MODE: Would upload file %s to DingDing API\n", filePath)
		return "test-media-id-12345", nil
	}

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create a multipart form body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file to the form
	part, err := writer.CreateFormFile("media", filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}

	// Copy the file content to the form
	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file content: %v", err)
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close multipart writer: %v", err)
	}

	// Construct the request URL
	requestURL := fmt.Sprintf("%s%s&type=file", bot.WebhookURL, bot.WebhookKey)
	
	// Add signature if sign key is provided
	if bot.SignKey != "" {
		timestamp := time.Now().UnixNano() / 1e6
		signature, err := bot.generateSignature(timestamp)
		if err != nil {
			return "", fmt.Errorf("failed to generate signature: %v", err)
		}
		requestURL = fmt.Sprintf("%s&timestamp=%d&sign=%s", requestURL, timestamp, url.QueryEscape(signature))
	}

	// Send the HTTP POST request
	resp, err := http.Post(requestURL, writer.FormDataContentType(), body)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	// Check for API errors
	if errcode, ok := result["errcode"].(float64); ok && errcode != 0 {
		errmsg, _ := result["errmsg"].(string)
		return "", fmt.Errorf("DingDing API error: %s", errmsg)
	}

	// Extract and return the media ID
	mediaID, ok := result["media_id"].(string)
	if !ok {
		return "", fmt.Errorf("media_id not found in response")
	}
	
	return mediaID, nil
}

// sendRequest sends a request to the DingDing API with the given payload.
// This is an internal helper method used by the public message sending methods.
// Parameters:
//   - payload: A map containing the message payload to send to the DingDing API
// Returns:
//   - An error if the request fails, nil otherwise
func (bot *DingDingBot) sendRequest(payload map[string]interface{}) error {
	// Convert the payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON payload: %v", err)
	}

	// Check if we're in test mode (webhook key starts with "test-")
	if len(bot.WebhookKey) >= 5 && bot.WebhookKey[:5] == "test-" {
		fmt.Printf("TEST MODE: Would send to DingDing API: %s\n", string(jsonPayload))
		return nil
	}

	// Construct the request URL
	requestURL := bot.WebhookURL + bot.WebhookKey
	
	// Add signature if sign key is provided
	if bot.SignKey != "" {
		timestamp := time.Now().UnixNano() / 1e6
		signature, err := bot.generateSignature(timestamp)
		if err != nil {
			return fmt.Errorf("failed to generate signature: %v", err)
		}
		requestURL = fmt.Sprintf("%s&timestamp=%d&sign=%s", requestURL, timestamp, url.QueryEscape(signature))
	}
	
	// Send the HTTP POST request
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	// Parse the response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	// Check for API errors
	if errcode, ok := result["errcode"].(float64); ok && errcode != 0 {
		errmsg, _ := result["errmsg"].(string)
		return fmt.Errorf("DingDing API error: %s", errmsg)
	}

	return nil
}
