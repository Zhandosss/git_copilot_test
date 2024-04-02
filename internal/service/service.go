package service

type OpenAi interface {
	GetAnsFromOpenAi(systemName, prompt string) (string, error)
	GeneratePrompts(systemName string) ([]string, error)
}

type Service struct {
	OpenAi
}

func New() *Service {
	return &Service{
		OpenAi: NewOpenAiService(),
	}
}
