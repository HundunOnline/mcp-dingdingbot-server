#!/bin/bash
echo "Testing send_news functionality..."
curl -X POST -H "Content-Type: application/json" -d '{
  "tool": "send_news",
  "params": {
    "title": "Test News",
    "text": "This is a test news message from DingDing Bot",
    "messageUrl": "https://github.com/HundunOnline",
    "picUrl": "https://avatars.githubusercontent.com/u/583231?v=4"
  }
}' http://localhost:8080/api/v1/run
