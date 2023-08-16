package repository

import "dinosaur-management/internal/domain"

type DinosaurRepository interface {
	ListDinosaurs() ([]domain.Dinosaur, error)
	CreateDinosaur(dino domain.Dinosaur) error
}
