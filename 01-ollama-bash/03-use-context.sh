#!/bin/bash

OLLAMA_URL="http://host.docker.internal:11434"
MODEL="phi3"

# First question

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

JSON_RESULT=$(curl -v ${OLLAMA_URL}/api/generate \
   -H "Content-Type: application/json" \
   -d "${DATA}" | jq -c '{ response, context }')

# Extract the value of "response" using jq
RESPONSE=$(echo "$JSON_RESULT" | jq -r '.response')

# Extract the value of "context" as a string using jq
CONTEXT=$(echo "$JSON_RESULT" | jq -r '.context')

echo "${RESPONSE}"
#echo "${CONTEXT}"

# Second question

read -r -d '' USER_CONTENT <<- EOM
Who is his best friend?
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
  "stream": false,
  "context": ${CONTEXT}
}
EOM

curl -v ${OLLAMA_URL}/api/generate \
   -H "Content-Type: application/json" \
   -d "${DATA}" | jq -c '{ response, context }'
