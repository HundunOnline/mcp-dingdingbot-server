#!/bin/bash
echo "Testing send_markdown functionality..."
curl -X POST -H "Content-Type: application/json" -d '{
  "tool": "send_markdown",
  "params": {
    "title": "Test Markdown",
    "content": "# This is a test markdown message\n## From DingDing Bot\n- Item 1\n- Item 2",
    "at_mobiles": [],
    "at_user_ids": [],
    "is_at_all": false
  }
}' http://localhost:8080/api/v1/run
