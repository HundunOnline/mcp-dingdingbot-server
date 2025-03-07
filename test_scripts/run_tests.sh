#!/bin/bash

# Start the server in the background with a test webhook key
export DINGDING_BOT_WEBHOOK_KEY="test-webhook-key"
cd ..
./dist/mcp-dingdingbot-server &
SERVER_PID=$!

# Wait for the server to start
sleep 2

# Run the test scripts
cd test_scripts
./test_text.sh
./test_markdown.sh
./test_image.sh
./test_news.sh
./test_template_card.sh

# Kill the server
kill $SERVER_PID

echo "All tests completed."
