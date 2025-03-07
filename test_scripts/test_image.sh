#!/bin/bash
echo "Testing send_image functionality..."
# Base64 encoded small test image
BASE64_IMAGE="iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8BQDwAEhQGAhKmMIQAAAABJRU5ErkJggg=="
MD5_HASH="5d41402abc4b2a76b9719d911017c592"

curl -X POST -H "Content-Type: application/json" -d "{
  \"tool\": \"send_image\",
  \"params\": {
    \"base64\": \"$BASE64_IMAGE\",
    \"md5\": \"$MD5_HASH\"
  }
}" http://localhost:8080/api/v1/run
