package usecase

import "dinosaur-management/internal/domain"

type DinosaurUsecase interface {
	ListDinosaurs() ([]domain.Dinosaur, error)
	CreateDinosaur(dino domain.Dinosaur) error
}
