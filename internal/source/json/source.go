package jsonsource

import (
	"encoding/json"
	"os"

	"sofiasoft/internal/domain"
)

type Source struct {
	filename string
}

func New(filename string) *Source {
	return &Source{
		filename: filename,
	}
}

func (s *Source) LoadPosts() ([]domain.Post, error) {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}

	var posts []domain.Post

	if err := json.Unmarshal(data, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
