package adapter

import "dinosaur-management/internal/domain"

type DinosaurAdapter interface {
	ListDinosaurs() ([]domain.Dinosaur, error)
	CreateDinosaur(dto DinosaurDTO) error
}

type DinosaurDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
