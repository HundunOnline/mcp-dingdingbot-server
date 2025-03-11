#!/bin/bash

# Example script to send a markdown message to DingDing group

# Set your webhook key and sign key (optional)
export DINGDING_BOT_WEBHOOK_KEY="your-webhook-key"
# export DINGDING_BOT_SIGN_KEY="your-sign-key"  # Uncomment if using signature verification

# JSON-RPC request to send a markdown message
echo '{
  "jsonrpc": "2.0",
  "id": "2",
  "method": "callTool",
  "params": {
    "tool": "send_markdown",
    "arguments": {
      "title": "Test Markdown",
      "content": "# This is a test markdown message\n## From DingDing Bot\n- Item 1\n- Item 2",
      "at_mobiles": "",
      "at_user_ids": "",
      "is_at_all": false
    }
  }
}' | mcp-dingdingbot-server
