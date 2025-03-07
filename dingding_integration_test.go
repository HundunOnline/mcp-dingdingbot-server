package main

import (
	"fmt"
	"testing"
)

func TestDingDingBotInTestMode(t *testing.T) {
	// Create a new DingDing bot with test webhook key
	bot := NewDingDingBot(DINGDING_BOT_SEND_URL, "test-webhook-key", "")

	// Test sending a text message
	err := bot.SendText("Test message", []string{}, []string{}, false)
	if err != nil {
		t.Errorf("Failed to send text message: %v", err)
	}

	// Test sending a markdown message
	err = bot.SendMarkdown("Test Title", "# Test markdown message", []string{}, []string{}, false)
	if err != nil {
		t.Errorf("Failed to send markdown message: %v", err)
	}

	// Test sending an image message
	err = bot.SendImage("SGVsbG8sIERpbmdEaW5nIQ==", "d41d8cd98f00b204e9800998ecf8427e")
	if err != nil {
		t.Errorf("Failed to send image message: %v", err)
	}

	// Test sending a news message
	err = bot.SendNews("Test News", "Test Description", "https://github.com/HundunOnline", "https://example.com/image.jpg")
	if err != nil {
		t.Errorf("Failed to send news message: %v", err)
	}

	// Test sending a template card message
	err = bot.SendTemplateCard("Test Card", "Test Card Content", "View Details", "https://github.com/HundunOnline", "0")
	if err != nil {
		t.Errorf("Failed to send template card message: %v", err)
	}

	// Test uploading a file (this will return a test media ID in test mode)
	mediaID, err := bot.UploadFile("test_file.txt")
	if err != nil {
		t.Errorf("Failed to upload file: %v", err)
	}
	if mediaID != "test-media-id-12345" {
		t.Errorf("Expected test media ID, got: %s", mediaID)
	}

	fmt.Println("All DingDing bot tests passed in test mode!")
}
