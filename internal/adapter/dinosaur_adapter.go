package adapter

import "dinosaur-management/internal/domain"

type DinosaurAdapter interface {
	ListDinosaurs() ([]domain.Dinosaur, error)
	CreateDinosaur(dto DinosaurDTO) error
}

type DinosaurDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
