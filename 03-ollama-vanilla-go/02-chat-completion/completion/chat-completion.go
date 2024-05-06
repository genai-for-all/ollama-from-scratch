package completion

import (
	"02-chat-completion/llm"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Chat(url string, query llm.Query) (llm.Answer, error) {

	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return llm.Answer{}, err
	}

	//fmt.Println(string(jsonQuery))

	req, err := http.NewRequest(http.MethodPost, url+"/api/chat", bytes.NewBuffer(jsonQuery))
	if err != nil {
		return llm.Answer{}, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return llm.Answer{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return llm.Answer{}, errors.New("Error: status code: " + resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return llm.Answer{}, err
	}
	//fmt.Println(string(body))

	var answer llm.Answer
	err = json.Unmarshal([]byte(string(body)), &answer)
	if err != nil {
		return llm.Answer{}, err
	}

	return answer, nil
}
