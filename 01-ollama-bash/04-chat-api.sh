#!/bin/bash

OLLAMA_URL="http://host.docker.internal:11434"
MODEL="deepseek-coder"

read -r -d '' SYSTEM_CONTENT <<- EOM
You are an expert in computer programming. 
Please make friendly answer for the noobs. 
Add source code examples if you can.
EOM

read -r -d '' USER_CONTENT <<- EOM
I need a clear explanation regarding the following question: 
Can you create a \"hello world\" program in Golang? 
And, please, be structured with bullet points.
EOM

SYSTEM_CONTENT=$(echo ${SYSTEM_CONTENT} | tr -d '\n')
USER_CONTENT=$(echo ${USER_CONTENT} | tr -d '\n')

read -r -d '' DATA <<- EOM
{
  "model":"${MODEL}",
  "options": {
    "temperature": 0.5,
    "repeat_last_n": 2
  },
  "messages": [
    {"role":"system", "content": "${SYSTEM_CONTENT}"},
    {"role":"user", "content": "${USER_CONTENT}"}
  ],
  "stream": false,
  "raw": false
}
EOM

curl -v ${OLLAMA_URL}/api/chat \
    -H "Content-Type: application/json" \
    -d "${DATA}" | jq '.message.content' > output.md
    
sed -i 's/\\n/\n/g' output.md
sed -i 's/\\"/"/g' output.md
sed -i 's/^"//' output.md



