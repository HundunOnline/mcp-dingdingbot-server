package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	webhookKey := os.Getenv("DINGDING_BOT_WEBHOOK_KEY")
	if webhookKey == "" {
		log.Println("DINGDING_BOT_WEBHOOK_KEY environment variable is required")
		return
	}

	// Get the sign key for signature verification (optional)
	signKey := os.Getenv("DINGDING_BOT_SIGN_KEY")

	bot := NewDingDingBot(DINGDING_BOT_SEND_URL, webhookKey, signKey)

	s := server.NewMCPServer(
		"mcp-dingdingbot-server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	sendTextTool := mcp.NewTool("send_text",
		mcp.WithDescription("Send a text message to DingDing group"),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Text content to send"),
		),
		mcp.WithString("at_mobiles",
			mcp.Description("List of mobile numbers to mention, multiple numbers use commas to separate, such as 13800138000,13800138001"),
		),
		mcp.WithString("at_user_ids",
			mcp.Description("List of user IDs to mention, multiple IDs use commas to separate"),
		),
		mcp.WithBoolean("is_at_all",
			mcp.Description("Whether to mention all users in the group"),
		),
	)
	s.AddTool(sendTextTool, sendTextHandler(bot))

	sendMarkdownTool := mcp.NewTool("send_markdown",
		mcp.WithDescription("Send a markdown message to DingDing group"),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Title of the markdown message"),
		),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Markdown content to send"),
		),
		mcp.WithString("at_mobiles",
			mcp.Description("List of mobile numbers to mention, multiple numbers use commas to separate"),
		),
		mcp.WithString("at_user_ids",
			mcp.Description("List of user IDs to mention, multiple IDs use commas to separate"),
		),
		mcp.WithBoolean("is_at_all",
			mcp.Description("Whether to mention all users in the group"),
		),
	)
	s.AddTool(sendMarkdownTool, sendMarkdownHandler(bot))

	sendImageTool := mcp.NewTool("send_image",
		mcp.WithDescription("Send an image message to DingDing group"),
		mcp.WithString("base64_data",
			mcp.Required(),
			mcp.Description("Base64 encoded image data"),
		),
		mcp.WithString("md5",
			mcp.Required(),
			mcp.Description("MD5 hash of the image"),
		),
	)
	s.AddTool(sendImageTool, sendImageHandler(bot))

	sendNewsTool := mcp.NewTool("send_news",
		mcp.WithDescription("Send a link message to DingDing group"),
		mcp.WithString("title", 
			mcp.Required(), 
			mcp.Description("Title of the link message")),
		mcp.WithString("text", 
			mcp.Required(), 
			mcp.Description("Text content of the link message")),
		mcp.WithString("message_url", 
			mcp.Required(), 
			mcp.Description("URL of the link message")),
		mcp.WithString("pic_url", 
			mcp.Description("Picture URL of the link message")),
	)
	s.AddTool(sendNewsTool, sendNewsHandler(bot))

	sendTemplateCardTool := mcp.NewTool("send_template_card",
		mcp.WithDescription("Send an action card message to DingDing group"),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Title of the action card"),
		),
		mcp.WithString("text",
			mcp.Required(),
			mcp.Description("Text content of the action card"),
		),
		mcp.WithString("single_title",
			mcp.Required(),
			mcp.Description("Title of the single button"),
		),
		mcp.WithString("single_url",
			mcp.Required(),
			mcp.Description("URL for the single button"),
		),
		mcp.WithString("btn_orientation",
			mcp.Description("Button orientation, 0: vertical, 1: horizontal"),
		),
	)
	s.AddTool(sendTemplateCardTool, sendTemplateCardHandler(bot))

	uploadFileTool := mcp.NewTool("upload_file",
		mcp.WithDescription("Upload a file to DingDing"),
		mcp.WithString("file_path",
			mcp.Required(),
			mcp.Description("Path to the file to upload"),
		),
	)
	s.AddTool(uploadFileTool, uploadFileHandler(bot))

	if err := server.ServeStdio(s); err != nil {
		log.Printf("Server error: %v\n", err)
	}
}

func sendTextHandler(bot *DingDingBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var atMobilesStr string
		var atUserIdsStr string
		var atMobiles []string
		var atUserIds []string
		var isAtAll bool

		content := request.Params.Arguments["content"].(string)

		if request.Params.Arguments["at_mobiles"] != nil {
			atMobilesStr = request.Params.Arguments["at_mobiles"].(string)
			atMobiles = strings.Split(atMobilesStr, ",")
		} else {
			atMobiles = []string{}
		}

		if request.Params.Arguments["at_user_ids"] != nil {
			atUserIdsStr = request.Params.Arguments["at_user_ids"].(string)
			atUserIds = strings.Split(atUserIdsStr, ",")
		} else {
			atUserIds = []string{}
		}

		if request.Params.Arguments["is_at_all"] != nil {
			isAtAll = request.Params.Arguments["is_at_all"].(bool)
		} else {
			isAtAll = false
		}

		err := bot.SendText(content, atMobiles, atUserIds, isAtAll)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to send text message: %v", err)), nil
		}

		return mcp.NewToolResultText("Text message sent successfully"), nil
	}
}

func sendMarkdownHandler(bot *DingDingBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		title := request.Params.Arguments["title"].(string)
		content := request.Params.Arguments["content"].(string)
		
		var atMobiles []string
		var atUserIds []string
		var isAtAll bool

		if request.Params.Arguments["at_mobiles"] != nil {
			atMobilesStr := request.Params.Arguments["at_mobiles"].(string)
			atMobiles = strings.Split(atMobilesStr, ",")
		} else {
			atMobiles = []string{}
		}

		if request.Params.Arguments["at_user_ids"] != nil {
			atUserIdsStr := request.Params.Arguments["at_user_ids"].(string)
			atUserIds = strings.Split(atUserIdsStr, ",")
		} else {
			atUserIds = []string{}
		}

		if request.Params.Arguments["is_at_all"] != nil {
			isAtAll = request.Params.Arguments["is_at_all"].(bool)
		} else {
			isAtAll = false
		}

		bot.WebhookURL = DINGDING_BOT_SEND_URL
		err := bot.SendMarkdown(title, content, atMobiles, atUserIds, isAtAll)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to send markdown message: %v", err)), nil
		}

		return mcp.NewToolResultText("Markdown message sent successfully"), nil
	}
}

func sendImageHandler(bot *DingDingBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		base64Data := request.Params.Arguments["base64_data"].(string)
		md5 := request.Params.Arguments["md5"].(string)

		bot.WebhookURL = DINGDING_BOT_SEND_URL
		err := bot.SendImage(base64Data, md5)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to send image message: %v", err)), nil
		}

		return mcp.NewToolResultText("Image message sent successfully"), nil
	}
}

func sendNewsHandler(bot *DingDingBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Extract the parameters from the request
		title := request.Params.Arguments["title"].(string)
		text := request.Params.Arguments["text"].(string)
		messageUrl := request.Params.Arguments["message_url"].(string)
		picUrl := ""
		
		if request.Params.Arguments["pic_url"] != nil {
			picUrl = request.Params.Arguments["pic_url"].(string)
		}

		// Send the news article
		bot.WebhookURL = DINGDING_BOT_SEND_URL
		err := bot.SendNews(title, text, messageUrl, picUrl)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to send news message: %v", err)), nil
		}

		// Return success result
		return mcp.NewToolResultText("News message sent successfully"), nil
	}
}

func sendTemplateCardHandler(bot *DingDingBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		title := request.Params.Arguments["title"].(string)
		text := request.Params.Arguments["text"].(string)
		singleTitle := request.Params.Arguments["single_title"].(string)
		singleURL := request.Params.Arguments["single_url"].(string)
		btnOrientation := "0"
		
		if request.Params.Arguments["btn_orientation"] != nil {
			btnOrientation = request.Params.Arguments["btn_orientation"].(string)
		}

		bot.WebhookURL = DINGDING_BOT_SEND_URL
		err := bot.SendTemplateCard(title, text, singleTitle, singleURL, btnOrientation)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to send template card message: %v", err)), nil
		}

		return mcp.NewToolResultText("Template card message sent successfully"), nil
	}
}

func uploadFileHandler(bot *DingDingBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		filePath := request.Params.Arguments["file_path"].(string)

		bot.WebhookURL = DINGDING_BOT_UPLOAD_URL
		mediaID, err := bot.UploadFile(filePath)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to upload file: %v", err)), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("File uploaded successfully, media ID: %s", mediaID)), nil
	}
}
