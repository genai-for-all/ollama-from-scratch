package completion

import (
	"05-embeddings-llm/llm"
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Chat(url string, query llm.Query) (llm.Answer, error) {

	query.Stream = false

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
	//err = json.Unmarshal([]byte(string(body)), &answer)
	err = json.Unmarshal(body, &answer)

	if err != nil {
		return llm.Answer{}, err
	}

	return answer, nil
}

func ChatStream(url string, query llm.Query, onChunk func(llm.Answer) error) error {
	query.Stream = true

	jsonQuery, err := json.Marshal(query)
	if err != nil {
		return err
	}

	resp, err := http.Post(url+"/api/chat", "application/json; charset=utf-8", bytes.NewBuffer(jsonQuery))
	if err != nil {
		return err
	}
	reader := bufio.NewReader(resp.Body)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		var answer llm.Answer
		err = json.Unmarshal(line, &answer)
		if err != nil {
			onChunk(llm.Answer{})
		}

		err = onChunk(answer)

		// generate an error to stop the stream
		if err != nil {
			return err
		}
	}

	return nil

}
