#!/bin/bash

#OLLAMA_URL="http://host.docker.internal:11434"
#OLLAMA_URL="http://localhost:11434"
OLLAMA_URL="http://robby.local:11434"

MODEL="all-minilm"

read -r -d '' DATA <<- EOM
{
  "name":"${MODEL}"
}
EOM

curl -v ${OLLAMA_URL}/api/pull \
   -H "Content-Type: application/json" \
   -d "${DATA}" | jq 


MODEL="all-minilm"

read -r -d '' DATA <<- EOM
{
  "name":"${MODEL}"
}
EOM

curl -v ${OLLAMA_URL}/api/pull \
   -H "Content-Type: application/json" \
   -d "${DATA}" | jq 
