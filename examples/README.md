# DingDing Bot Examples

This directory contains examples of how to use the mcp-dingdingbot-server to send various types of messages to DingDing group robots.

## Prerequisites

Before running these examples, make sure you have:

1. Installed the mcp-dingdingbot-server
2. Configured your DingDing bot webhook key and sign key (if using signature verification)

## Examples

### Text Message

```json
{
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
}
```

### Markdown Message

```json
{
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
}
```

### Image Message

```json
{
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
}
```

### News Message

```json
{
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
}
```

### Template Card Message

```json
{
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
}
```

## Running the Examples

You can run these examples by piping the JSON to the mcp-dingdingbot-server:

```bash
echo '{JSON_CONTENT}' | mcp-dingdingbot-server
```

Replace `{JSON_CONTENT}` with the JSON example you want to run.

## Environment Variables

Make sure to set the required environment variables:

```bash
export DINGDING_BOT_WEBHOOK_KEY="your-webhook-key"
export DINGDING_BOT_SIGN_KEY="your-sign-key"  # Optional
```

## Testing Script

You can also use the test script provided in the `test_scripts` directory:

```bash
cd test_scripts
./test_mcp.sh
```

This script will run all the examples and show the results.
