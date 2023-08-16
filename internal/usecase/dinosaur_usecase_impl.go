package usecase

import "dinosaur-management/internal/domain"

type DinosaurUsecaseImpl struct {
	dinoRepo domain.DinosaurRepository
}

func NewDinosaurUsecase(dinoRepo domain.DinosaurRepository) *DinosaurUsecaseImpl {
	return &DinosaurUsecaseImpl{dinoRepo: dinoRepo}
}

func (uc *DinosaurUsecaseImpl) ListDinosaurs() ([]domain.Dinosaur, error) {
	return uc.dinoRepo.ListDinosaurs()
}

func (uc *DinosaurUsecaseImpl) CreateDinosaur(dino domain.Dinosaur) error {
	return uc.dinoRepo.CreateDinosaur(dino)
}
