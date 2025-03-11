#!/bin/bash

# Example script to send a text message to DingDing group

# Set your webhook key and sign key (optional)
export DINGDING_BOT_WEBHOOK_KEY="your-webhook-key"
# export DINGDING_BOT_SIGN_KEY="your-sign-key"  # Uncomment if using signature verification

# JSON-RPC request to send a text message
echo '{
  "jsonrpc": "2.0",
  "id": "1",
  "method": "callTool",
  "params": {
    "tool": "send_text",
    "arguments": {
      "content": "This is a test message from DingDing Bot",
      "at_mobiles": "",
      "at_user_ids": "",
      "is_at_all": false
    }
  }
}' | mcp-dingdingbot-server
