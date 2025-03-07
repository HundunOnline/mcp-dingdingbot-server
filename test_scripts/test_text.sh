#!/bin/bash
echo "Testing send_text functionality..."
curl -X POST -H "Content-Type: application/json" -d '{
  "tool": "send_text",
  "params": {
    "content": "This is a test message from DingDing Bot",
    "at_mobiles": [],
    "at_user_ids": [],
    "is_at_all": false
  }
}' http://localhost:8080/api/v1/run
