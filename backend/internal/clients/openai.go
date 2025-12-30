package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type OpenAIClient struct {
	Endpoint   string
	APIKey     string
	Deployment string
}

func NewOpenAIClient(endpoint, key, deployment string) (*OpenAIClient, error) {
	if endpoint == "" || key == "" || deployment == "" {
		return nil, errors.New("azure openai config missing")
	}

	return &OpenAIClient{
		Endpoint:   endpoint,
		APIKey:     key,
		Deployment: deployment,
	}, nil
}

type chatRequest struct {
	Messages []map[string]string `json:"messages"`
}

type chatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (c *OpenAIClient) Chat(systemPrompt, userPrompt string) (string, error) {
	body := chatRequest{
		Messages: []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": userPrompt},
		},
	}

	b, _ := json.Marshal(body)

	req, _ := http.NewRequest(
		"POST",
		c.Endpoint+"/openai/deployments/"+c.Deployment+"/chat/completions?api-version=2024-02-15-preview", bytes.NewBuffer(b),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", c.APIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("azure openai request failed")
	}

	var parsed chatResponse
	json.NewDecoder(resp.Body).Decode(&parsed)

	if len(parsed.Choices) == 0 {
		return "", errors.New("empty response fro azure openai")
	}

	return parsed.Choices[0].Message.Content, nil
}
