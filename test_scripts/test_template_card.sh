#!/bin/bash
echo "Testing send_template_card functionality..."
curl -X POST -H "Content-Type: application/json" -d '{
  "tool": "send_template_card",
  "params": {
    "title": "Test Template Card",
    "text": "This is a test template card message from DingDing Bot",
    "singleTitle": "View Details",
    "singleURL": "https://github.com/HundunOnline",
    "btnOrientation": "0"
  }
}' http://localhost:8080/api/v1/run
