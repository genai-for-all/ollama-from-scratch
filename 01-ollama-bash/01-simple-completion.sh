#!/bin/bash

OLLAMA_URL="http://host.docker.internal:11434"
MODEL="phi3"

read -r -d '' USER_CONTENT <<- EOM
Who is James T Kirk?
EOM

USER_CONTENT=$(echo ${USER_CONTENT} | tr -d '\n')

read -r -d '' DATA <<- EOM
{
  "model":"${MODEL}",
  "options": {
    "temperature": 0.5,
    "repeat_last_n": 2
  },
  "prompt": "${USER_CONTENT}",
  "stream": false
}
EOM

curl -v ${OLLAMA_URL}/api/generate \
   -H "Content-Type: application/json" \
   -d "${DATA}" | jq -c '{ response, context }'



