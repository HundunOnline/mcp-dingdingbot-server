package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// MockWeComServer is a mock server for testing WeComBot.
type MockWeComServer struct {
	*httptest.Server
}

// test upload file need to setting your's test key
var testkey = "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx"

// NewMockWeComServer creates a new mock WeCom server.
func NewMockWeComServer() *MockWeComServer {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, _ := io.ReadAll(r.Body)
			fmt.Println(string(body))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"errcode": 0,
				"errmsg":  "ok",
			})
		}
	})

	server := httptest.NewServer(handler)
	return &MockWeComServer{Server: server}
}

// TestSendText tests the SendText method.
func TestSendText(t *testing.T) {
	mockServer := NewMockWeComServer()
	defer mockServer.Close()

	bot := NewWeComBot(mockServer.URL, "")
	err := bot.SendText("Hello, WeCom!", []string{}, []string{})
	if err != nil {
		t.Errorf("SendText failed: %v", err)
	}
}

// TestSendMarkdown tests the SendMarkdown method.
func TestSendMarkdown(t *testing.T) {
	mockServer := NewMockWeComServer()
	defer mockServer.Close()

	bot := NewWeComBot(mockServer.URL, "")
	err := bot.SendMarkdown("## Hello, WeCom!\nThis is a markdown message.")
	if err != nil {
		t.Errorf("SendMarkdown failed: %v", err)
	}
}

// TestSendImage tests the SendImage method.
func TestSendImage(t *testing.T) {
	mockServer := NewMockWeComServer()
	defer mockServer.Close()

	bot := NewWeComBot(mockServer.URL, "")
	err := bot.SendImage("SGVsbG8sIFdlQ29tIQ==", "d41d8cd98f00b204e9800998ecf8427e")
	if err != nil {
		t.Errorf("SendImage failed: %v", err)
	}
}

// TestSendNews tests the SendNews method.
func TestSendNews(t *testing.T) {
	mockServer := NewMockWeComServer()
	defer mockServer.Close()

	bot := NewWeComBot(mockServer.URL, "")
	articles := []NewsArticle{
		{
			Title:       "News Title",
			Description: "News Description",
			URL:         "https://example.com",
			PicURL:      "https://example.com/image.jpg",
		},
	}
	err := bot.SendNews(articles)
	if err != nil {
		t.Errorf("SendNews failed: %v", err)
	}
}

// TestSendTemplateCard tests the SendTemplateCard method.
func TestSendTemplateCard(t *testing.T) {
	mockServer := NewMockWeComServer()
	defer mockServer.Close()

	bot := NewWeComBot(mockServer.URL, "")
	err := bot.SendTemplateCard("text_notice", "Main Title", "Main Description", 1, "https://example.com", "", "")
	if err != nil {
		t.Errorf("SendTemplateCard failed: %v", err)
	}
}

// TestUploadFile tests the UploadFile method.
func TestUploadFile(t *testing.T) {
	mockServer := NewMockWeComServer()
	defer mockServer.Close()

	mockServer.URL = WECOM_BOT_UPLOAD_URL
	bot := NewWeComBot(mockServer.URL, testkey)

	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write some content to the temporary file
	_, err = tempFile.WriteString("Hello, WeCom!")
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tempFile.Close()

	mediaID, err := bot.UploadFile(tempFile.Name())
	if err != nil {
		t.Errorf("UploadFile failed: %v", err)
	} else {
		t.Logf("UploadFile succeeded, media ID: %s", mediaID)
	}
}
