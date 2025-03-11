#!/bin/bash

# Example script to send an image message to DingDing group

# Set your webhook key and sign key (optional)
export DINGDING_BOT_WEBHOOK_KEY="your-webhook-key"
# export DINGDING_BOT_SIGN_KEY="your-sign-key"  # Uncomment if using signature verification

# JSON-RPC request to send an image message
echo '{
  "jsonrpc": "2.0",
  "id": "3",
  "method": "callTool",
  "params": {
    "tool": "send_image",
    "arguments": {
      "base64_data": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg==",
      "md5": "5d41402abc4b2a76b9719d911017c592"
    }
  }
}' | mcp-dingdingbot-server
