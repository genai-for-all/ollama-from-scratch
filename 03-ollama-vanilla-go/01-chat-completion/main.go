package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var ollamaUrl = "http://host.docker.internal:11434"
var model = "deepseek-coder"
//var model = "phi3"

func main() {

	systemContent := `You are an expert in computer programming.
	Please make friendly answer for the noobs.
	Add source code examples if you can.`

	userContent := `I need a clear explanation regarding the following question:
	Can you create a "hello world" program in Golang?
	And, please, be structured with bullet points`

	messages := []map[string]string{
		{"role": "system", "content": systemContent},
		{"role": "user", "content": userContent},
	}

	data := map[string]interface{}{
		"model":    model,
		"messages": messages,
		"options": map[string]interface{}{
			"temperature":   0.5,
			"repeat_last_n": 2,
		},
		"stream": false,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("😡 Error marshalling JSON:", err)
	}

	req, err := http.NewRequest(http.MethodPost, ollamaUrl+"/api/chat", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("😡 Error creating request:", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("😡 Error sending request:", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("😡 Error:", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("😡 Error reading response body:", err)
	}
	//fmt.Println(string(body))

	//var answer map[string]interface{}
	answer := map[string]any{}
	err = json.Unmarshal([]byte(string(body)), &answer)
	if err != nil {
		fmt.Println("😡 Error unmarshalling JSON:", err)
	}
	fmt.Println(answer["message"].(map[string]any)["content"])

}
