#!/bin/bash

# Test script for mcp-dingdingbot-server
# This script sends test commands to the MCP server via stdin and captures the output

export DINGDING_BOT_WEBHOOK_KEY="test-webhook-key"
cd ..

echo "Starting MCP server in test mode..."
echo "Testing send_text functionality..."
echo '{"jsonrpc":"2.0","id":"1","method":"callTool","params":{"tool":"send_text","arguments":{"content":"This is a test message from DingDing Bot","at_mobiles":"","at_user_ids":"","is_at_all":false}}}' | ./dist/mcp-dingdingbot-server

echo "Testing send_markdown functionality..."
echo '{"jsonrpc":"2.0","id":"2","method":"callTool","params":{"tool":"send_markdown","arguments":{"title":"Test Markdown","content":"# This is a test markdown message\n## From DingDing Bot\n- Item 1\n- Item 2","at_mobiles":"","at_user_ids":"","is_at_all":false}}}' | ./dist/mcp-dingdingbot-server

echo "Testing send_image functionality..."
echo '{"jsonrpc":"2.0","id":"3","method":"callTool","params":{"tool":"send_image","arguments":{"base64_data":"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg==","md5":"5d41402abc4b2a76b9719d911017c592"}}}' | ./dist/mcp-dingdingbot-server

echo "Testing send_news functionality..."
echo '{"jsonrpc":"2.0","id":"4","method":"callTool","params":{"tool":"send_news","arguments":{"title":"Test News","text":"This is a test news message from DingDing Bot","message_url":"https://github.com/HundunOnline","pic_url":"https://avatars.githubusercontent.com/u/583231?v=4"}}}' | ./dist/mcp-dingdingbot-server

echo "Testing send_template_card functionality..."
echo '{"jsonrpc":"2.0","id":"5","method":"callTool","params":{"tool":"send_template_card","arguments":{"title":"Test Template Card","text":"This is a test template card message from DingDing Bot","single_title":"View Details","single_url":"https://github.com/HundunOnline","btn_orientation":"0"}}}' | ./dist/mcp-dingdingbot-server

echo "All tests completed."
