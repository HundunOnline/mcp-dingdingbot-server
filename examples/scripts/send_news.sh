#!/bin/bash

# Example script to send a news message to DingDing group

# Set your webhook key and sign key (optional)
export DINGDING_BOT_WEBHOOK_KEY="your-webhook-key"
# export DINGDING_BOT_SIGN_KEY="your-sign-key"  # Uncomment if using signature verification

# JSON-RPC request to send a news message
echo '{
  "jsonrpc": "2.0",
  "id": "4",
  "method": "callTool",
  "params": {
    "tool": "send_news",
    "arguments": {
      "title": "Test News",
      "text": "This is a test news message from DingDing Bot",
      "message_url": "https://github.com/HundunOnline",
      "pic_url": "https://avatars.githubusercontent.com/u/583231?v=4"
    }
  }
}' | mcp-dingdingbot-server
