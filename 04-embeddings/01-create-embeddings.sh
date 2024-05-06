#!/bin/bash

#OLLAMA_URL="http://host.docker.internal:11434"
OLLAMA_URL="http://localhost:11434"

MODEL="all-minilm"

read -r -d '' PROMPT <<- EOM
Who is James T Kirk?
EOM

#PROMPT=$(echo ${PROMPT} | tr -d '\n')

read -r -d '' DATA <<- EOM
{
  "model":"${MODEL}",
  "prompt": "${PROMPT}"
}
EOM

echo "${DATA}"

curl -v ${OLLAMA_URL}/api/embeddings \
   -H "Content-Type: application/json" \
   -d "${DATA}" | jq -c '{ embedding }'



