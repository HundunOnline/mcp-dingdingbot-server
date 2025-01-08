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
	webhookKey := os.Getenv("WECOM_BOT_WEBHOOK_KEY")
	if webhookKey == "" {
		log.Println("WECOM_BOT_WEBHOOK_KEY environment variable is required")
		return
	}

	bot := NewWeComBot(WECOM_BOT_SEND_URL, webhookKey)

	s := server.NewMCPServer(
		"mcp-wecombot-server",
		"1.0.0",
		server.WithResourceCapabilities(true, true),
		server.WithLogging(),
	)

	sendTextTool := mcp.NewTool("send_text",
		mcp.WithDescription("Send a text message to WeCom group"),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Text content to send"),
		),
		mcp.WithString("mentioned_list",
			mcp.Description("List of user IDs to mention,Multiple people use commas to separate, such as @xiaoyang, @wike."),
		),
		mcp.WithString("mentioned_mobile_list",
			mcp.Description("List of mobile numbers to mention,Multiple people use commas to separate, such as @xiaoyang, @wike."),
		),
	)
	s.AddTool(sendTextTool, sendTextHandler(bot))

	sendMarkdownTool := mcp.NewTool("send_markdown",
		mcp.WithDescription("Send a markdown message to WeCom group"),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Markdown content to send"),
		),
	)
	s.AddTool(sendMarkdownTool, sendMarkdownHandler(bot))

	sendImageTool := mcp.NewTool("send_image",
		mcp.WithDescription("Send an image message to WeCom group"),
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
		mcp.WithDescription("Send a news message to WeCom group"),
		mcp.WithString("title", mcp.Required(), mcp.Description("Title of the news article")),
		mcp.WithString("description", mcp.Description("Description of the news article")),
		mcp.WithString("url", mcp.Required(), mcp.Description("URL of the news article")),
		mcp.WithString("picurl", mcp.Description("Picture URL of the news article")),
	)
	s.AddTool(sendNewsTool, sendNewsHandler(bot))

	sendTemplateCardTool := mcp.NewTool("send_template_card",
		mcp.WithDescription("Send a template card message to WeCom group"),
		mcp.WithString("card_type",
			mcp.Required(),
			mcp.Description("Type of the template card"),
		),
		mcp.WithString("main_title",
			mcp.Required(),
			mcp.Description("Main title of the template card"),
		),
		mcp.WithString("main_desc",
			mcp.Description("Main description of the template card"),
		),
		mcp.WithNumber("card_action_type",
			mcp.Required(),
			mcp.Description("Type of the card action"),
		),
		mcp.WithString("card_action_url",
			mcp.Description("URL for the card action"),
		),
		mcp.WithString("card_action_appid",
			mcp.Description("App ID for the card action"),
		),
		mcp.WithString("card_action_pagepath",
			mcp.Description("Page path for the card action"),
		),
	)
	s.AddTool(sendTemplateCardTool, sendTemplateCardHandler(bot))

	uploadFileTool := mcp.NewTool("upload_file",
		mcp.WithDescription("Upload a file to WeCom"),
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
func sendTextHandler(bot *WeComBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		var mentionedListStr string
		var mentionedMobileListStr string
		var mentionedList []string
		var mentionedMobileList []string

		content := request.Params.Arguments["content"].(string)

		if request.Params.Arguments["mentioned_list"] == nil && request.Params.Arguments["mentioned_mobile_list"] == nil {
			mentionedListStr = ""
			mentionedMobileListStr = ""
		} else {
			mentionedListStr = request.Params.Arguments["mentioned_list"].(string)
			mentionedMobileListStr = request.Params.Arguments["mentioned_mobile_list"].(string)
		}

		if mentionedListStr != "" && mentionedMobileListStr != "" {
			mentionedList = strings.Split(mentionedListStr, ",")
			mentionedMobileList = strings.Split(mentionedMobileListStr, ",")
		} else {
			mentionedList = []string{}
			mentionedMobileList = []string{}
		}

		err := bot.SendText(content, mentionedList, mentionedMobileList)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to send text message: %v", err)), nil
		}

		return mcp.NewToolResultText("Text message sent successfully"), nil
	}
}

func sendMarkdownHandler(bot *WeComBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		content := request.Params.Arguments["content"].(string)

		bot.WebhookURL = WECOM_BOT_SEND_URL
		err := bot.SendMarkdown(content)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to send markdown message: %v", err)), nil
		}

		return mcp.NewToolResultText("Markdown message sent successfully"), nil
	}
}

func sendImageHandler(bot *WeComBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		base64Data := request.Params.Arguments["base64_data"].(string)
		md5 := request.Params.Arguments["md5"].(string)

		bot.WebhookURL = WECOM_BOT_SEND_URL
		err := bot.SendImage(base64Data, md5)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to send image message: %v", err)), nil
		}

		return mcp.NewToolResultText("Image message sent successfully"), nil
	}
}

func sendNewsHandler(bot *WeComBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// Extract the parameters from the request
		title := request.Params.Arguments["title"].(string)
		description := request.Params.Arguments["description"].(string)
		url := request.Params.Arguments["url"].(string)
		picurl := request.Params.Arguments["picurl"].(string)

		// Create a single NewsArticle
		newsArticle := NewsArticle{
			Title:       title,
			Description: description,
			URL:         url,
			PicURL:      picurl,
		}

		// Send the news article
		bot.WebhookURL = WECOM_BOT_SEND_URL
		err := bot.SendNews([]NewsArticle{newsArticle})
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to send news message: %v", err)), nil
		}

		// Return success result
		return mcp.NewToolResultText("News message sent successfully"), nil
	}
}

func sendTemplateCardHandler(bot *WeComBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		cardType := request.Params.Arguments["card_type"].(string)
		mainTitle := request.Params.Arguments["main_title"].(string)
		mainDesc := request.Params.Arguments["main_desc"].(string)
		cardActionType := request.Params.Arguments["card_action_type"].(int)
		cardActionURL := request.Params.Arguments["card_action_url"].(string)
		cardActionAppID := request.Params.Arguments["card_action_appid"].(string)
		cardActionPagePath := request.Params.Arguments["card_action_pagepath"].(string)

		bot.WebhookURL = WECOM_BOT_SEND_URL
		err := bot.SendTemplateCard(cardType, mainTitle, mainDesc, cardActionType, cardActionURL, cardActionAppID, cardActionPagePath)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to send template card message: %v", err)), nil
		}

		return mcp.NewToolResultText("Template card message sent successfully"), nil
	}
}

func uploadFileHandler(bot *WeComBot) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		filePath := request.Params.Arguments["file_path"].(string)

		bot.WebhookURL = WECOM_BOT_UPLOAD_URL
		mediaID, err := bot.UploadFile(filePath)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to upload file: %v", err)), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("File uploaded successfully, media ID: %s", mediaID)), nil
	}
}
