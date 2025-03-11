#!/bin/bash

# Example script to upload a file to DingDing

# Set your webhook key and sign key (optional)
export DINGDING_BOT_WEBHOOK_KEY="your-webhook-key"
# export DINGDING_BOT_SIGN_KEY="your-sign-key"  # Uncomment if using signature verification

# Create a temporary file to upload
echo "This is a test file for DingDing Bot" > test_file.txt

# JSON-RPC request to upload a file
echo '{
  "jsonrpc": "2.0",
  "id": "6",
  "method": "callTool",
  "params": {
    "tool": "upload_file",
    "arguments": {
      "file_path": "test_file.txt"
    }
  }
}' | mcp-dingdingbot-server

# Clean up the temporary file
rm test_file.txt
