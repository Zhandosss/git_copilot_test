package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git_copilot_test/internal/model"
	"git_copilot_test/internal/service/prompts"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	ErrPromptTemplateIsEmpty = fmt.Errorf("prompt template is empty")
)

type OpenAiService struct {
	promptsTemplate []string
	apiKey          string
	client          *http.Client
}

func NewOpenAiService() *OpenAiService {
	return &OpenAiService{
		promptsTemplate: prompts.Generate(),
		apiKey:          os.Getenv("OPENAI_KEY"),
		client:          &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *OpenAiService) GetAnsFromOpenAi(systemName, prompt string) (string, error) {
	reqBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": "In this conversation, you provide information about the API of a provider. The provider is called " + systemName,
			},
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"temperature": 0.0,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code is not 200: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result model.InfoResponse
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return "", err
	}

	return result.Choices[0].Msg.Content, nil
}

func (s *OpenAiService) GeneratePrompts(systemName string) ([]string, error) {
	if len(s.promptsTemplate) == 0 {
		return nil, ErrPromptTemplateIsEmpty
	}
	providerPrompts := make([]string, len(s.promptsTemplate))
	for i, prompt := range s.promptsTemplate {
		providerPrompts[i] = fmt.Sprintf(prompt, systemName)
	}
	return providerPrompts, nil
}
