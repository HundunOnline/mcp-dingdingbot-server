#!/bin/bash

# Example script to send a template card message to DingDing group

# Set your webhook key and sign key (optional)
export DINGDING_BOT_WEBHOOK_KEY="your-webhook-key"
# export DINGDING_BOT_SIGN_KEY="your-sign-key"  # Uncomment if using signature verification

# JSON-RPC request to send a template card message
echo '{
  "jsonrpc": "2.0",
  "id": "5",
  "method": "callTool",
  "params": {
    "tool": "send_template_card",
    "arguments": {
      "title": "Test Template Card",
      "text": "This is a test template card message from DingDing Bot",
      "single_title": "View Details",
      "single_url": "https://github.com/HundunOnline",
      "btn_orientation": "0"
    }
  }
}' | mcp-dingdingbot-server
