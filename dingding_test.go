package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockDingDingServer is a mock server for testing DingDingBot.
type MockDingDingServer struct {
	*httptest.Server
}

// test upload file need to setting your's test key
var testkey = "xxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx"

// NewMockDingDingServer creates a new mock DingDing server.
func NewMockDingDingServer() *MockDingDingServer {
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
	return &MockDingDingServer{Server: server}
}

// TestSendText tests the SendText method.
func TestSendText(t *testing.T) {
	mockServer := NewMockDingDingServer()
	defer mockServer.Close()

	bot := NewDingDingBot(mockServer.URL, "")
	err := bot.SendText("Hello, DingDing!", []string{}, []string{}, false)
	if err != nil {
		t.Errorf("SendText failed: %v", err)
	}
}

// TestSendMarkdown tests the SendMarkdown method.
func TestSendMarkdown(t *testing.T) {
	mockServer := NewMockDingDingServer()
	defer mockServer.Close()

	bot := NewDingDingBot(mockServer.URL, "")
	err := bot.SendMarkdown("Hello", "## Hello, DingDing!\nThis is a markdown message.", []string{}, []string{}, false)
	if err != nil {
		t.Errorf("SendMarkdown failed: %v", err)
	}
}

// TestSendImage tests the SendImage method.
func TestSendImage(t *testing.T) {
	mockServer := NewMockDingDingServer()
	defer mockServer.Close()

	bot := NewDingDingBot(mockServer.URL, "")
	err := bot.SendImage("SGVsbG8sIERpbmdEaW5nIQ==", "d41d8cd98f00b204e9800998ecf8427e")
	if err != nil {
		t.Errorf("SendImage failed: %v", err)
	}
}

// TestSendNews tests the SendNews method.
func TestSendNews(t *testing.T) {
	mockServer := NewMockDingDingServer()
	defer mockServer.Close()

	bot := NewDingDingBot(mockServer.URL, "")
	err := bot.SendNews("News Title", "News Description", "https://example.com", "https://example.com/image.jpg")
	if err != nil {
		t.Errorf("SendNews failed: %v", err)
	}
}

// TestSendTemplateCard tests the SendTemplateCard method.
func TestSendTemplateCard(t *testing.T) {
	mockServer := NewMockDingDingServer()
	defer mockServer.Close()

	bot := NewDingDingBot(mockServer.URL, "")
	err := bot.SendTemplateCard("Main Title", "Main Description", "View Details", "https://example.com", "0")
	if err != nil {
		t.Errorf("SendTemplateCard failed: %v", err)
	}
}

// TestUploadFile tests the UploadFile method.
func TestUploadFile(t *testing.T) {
	// For this test, we'll just skip the actual upload since it requires a real API
	// In a real environment, this would be tested with proper mocks or integration tests
	t.Skip("Skipping upload file test as it requires a real DingDing API")
}
