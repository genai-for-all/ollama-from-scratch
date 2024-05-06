package main

import (
	"05-embeddings-llm/llm"
	"05-embeddings-llm/completion"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

//var ollamaUrl = "http://host.docker.internal:11434"

var ollamaUrl = "http://localhost:11434"

// var embeddingsModel = "deepseek-coder"
// var embeddingsModel = "phi3"
var embeddingsModel = "all-minilm" // This model is for the embeddings of the documents
var chatModel = "qwen:0.5b"        // This model is for the chat completion

type EmbeddingResponse struct {
	Embedding []float64 `json:"embedding"`
}

type VectorRecord struct {
	Id        string
	Prompt    string
	Embedding []float64
}

type Query4Embedding struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
}

func CreateEmbedding(ollamaUrl string, query Query4Embedding, id string) (VectorRecord, error) {
	jsonData, err := json.Marshal(query)
	if err != nil {
		log.Fatal("😡 Error marshalling JSON:", err)
		return VectorRecord{}, err
	}

	req, err := http.NewRequest(http.MethodPost, ollamaUrl+"/api/embeddings", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("😡 Error creating request:", err)
		return VectorRecord{}, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("😡 Error sending request:", err)
		return VectorRecord{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("😡 Error:", resp.StatusCode)
		return VectorRecord{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("😡 Error reading response body:", err)
		return VectorRecord{}, err
	}

	var answer EmbeddingResponse
	err = json.Unmarshal([]byte(string(body)), &answer)
	if err != nil {
		fmt.Println("😡 Error unmarshalling JSON:", err)
		return VectorRecord{}, err
	}

	vectorRecord := VectorRecord{
		Prompt:    query.Prompt,
		Embedding: answer.Embedding,
		Id:        id,
	}

	return vectorRecord, nil
}

func dotProduct(v1 []float64, v2 []float64) float64 {
	// Calculate the dot product of two vectors
	sum := 0.0
	for i := range v1 {
		sum += v1[i] * v2[i]
	}
	return sum
}

func CosineDistance(v1, v2 []float64) float64 {
	// Calculate the cosine distance between two vectors
	product := dotProduct(v1, v2)
	norm1 := math.Sqrt(dotProduct(v1, v1))
	norm2 := math.Sqrt(dotProduct(v2, v2))
	if norm1 <= 0.0 || norm2 <= 0.0 {
		// Handle potential division by zero
		return 0.0
	}
	return product / (norm1 * norm2)
}

//! naive implementation of embedding similarity

/*
	If I split the documents into smaller pieces, how I will find the relation ship between the vectors?
	one way: (an idea)
	create an embedding from each document
	split a doc into smaller pieces -> create records with the same embedding
	or create a record: relatedSplit?
*/

func main() {

	docs := []string{
		`Michael Burnham is the main character on the Star Trek series, Discovery.  
		She's a human raised on the logical planet Vulcan by Spock's father.  
		Burnham is intelligent and struggles to balance her human emotions with Vulcan logic.  
		She's become a Starfleet captain known for her determination and problem-solving skills.
		Originally played by actress Sonequa Martin-Green`,

		`James T. Kirk, also known as Captain Kirk, is a fictional character from the Star Trek franchise.  
		He's the iconic captain of the starship USS Enterprise, 
		boldly exploring the galaxy with his crew.  
		Originally played by actor William Shatner, 
		Kirk has appeared in TV series, movies, and other media.`,

		`Jean-Luc Picard is a fictional character in the Star Trek franchise.
		He's most famous for being the captain of the USS Enterprise-D,
		a starship exploring the galaxy in the 24th century.
		Picard is known for his diplomacy, intelligence, and strong moral compass.
		He's been portrayed by actor Patrick Stewart.`,

		`Lieutenant Philippe Charrière, known as the **Silent Sentinel** of the USS Discovery, 
		is the enigmatic programming genius whose codes safeguard the ship's secrets and operations. 
		His swift problem-solving skills are as legendary as the mysterious aura that surrounds him. 
		Charrière, a man of few words, speaks the language of machines with unrivaled fluency, 
		making him the crew's unsung guardian in the cosmos. His best friend is Spiderman from the Marvel Cinematic Universe.`,
	}

	v1, _ := CreateEmbedding(ollamaUrl, Query4Embedding{Prompt: docs[0], Model: embeddingsModel}, "Michael Burnham")
	v2, _ := CreateEmbedding(ollamaUrl, Query4Embedding{Prompt: docs[1], Model: embeddingsModel}, "James T. Kirk")
	v3, _ := CreateEmbedding(ollamaUrl, Query4Embedding{Prompt: docs[2], Model: embeddingsModel}, "Jean-Luc Picard")
	v4, _ := CreateEmbedding(ollamaUrl, Query4Embedding{Prompt: docs[3], Model: embeddingsModel}, "Philippe Charrière")

	vectorsList := []VectorRecord{v1, v2, v3, v4}


	// a message of the Chat system
	//userContent := `Who is Philippe Charrière and what spaceship does he work on?`
	userContent := `What is the nickname of Philippe Charrière?`



	// Create an embedding from a question
	embeddingFromQuestion, _ := CreateEmbedding(
		ollamaUrl, Query4Embedding{
			Prompt: userContent,
			Model:  embeddingsModel},
		"question-1",
	)

	fmt.Println(embeddingFromQuestion.Prompt, ":")
	var maxDistance float64 = 0.0
	var selectedIdx int
	for idx, v := range vectorsList {
		distance := CosineDistance(embeddingFromQuestion.Embedding, v.Embedding)
		if distance > maxDistance {
			maxDistance = distance
			selectedIdx = idx
		}
		fmt.Println("  - ", idx, v.Id, distance)
	}
	fmt.Println("Selected:", vectorsList[selectedIdx].Prompt)
	// I take only the nearest vector to the question

	systemContent := `    You are an AI assistant. Your name is Seven. 
    Some people are calling you Seven of Nine.
    You are an expert in Star Trek.
    All questions are about Star Trek.
    Using the provided context, answer the user's question
    to the best of your ability using only the resources provided.`

	documentsContent := `<context><doc>` + vectorsList[selectedIdx].Prompt + `</doc></context>`


	query := llm.Query{
		Model: chatModel,
		Messages: []llm.Message{
			{Role: "system", Content: systemContent},
			{Role: "system", Content: documentsContent},
			{Role: "user", Content: userContent},
		},
		Options: llm.Options{
			Temperature: 0.5,
			RepeatLastN: 2,
		},
		Stream: false,
	}

	err := completion.ChatStream(ollamaUrl, query,
		func(answer llm.Answer) error {
			fmt.Print(answer.Message.Content)
			return nil
		})

	if err != nil {
		log.Fatal("😡:", err)
	}

}
