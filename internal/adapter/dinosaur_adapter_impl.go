package adapter

import (
	"dinosaur-management/internal/domain"
	"dinosaur-management/internal/usecase"
)

type DinosaurAdapterImpl struct {
	dinoUsecase usecase.DinosaurUsecase
}

func NewDinosaurAdapter(dinoUsecase usecase.DinosaurUsecase) *DinosaurAdapterImpl {
	return &DinosaurAdapterImpl{dinoUsecase: dinoUsecase}
}

func (a *DinosaurAdapterImpl) ListDinosaurs() ([]domain.Dinosaur, error) {
	return a.dinoUsecase.ListDinosaurs()
}

func (a *DinosaurAdapterImpl) CreateDinosaur(dto DinosaurDTO) error {
	newDinosaur := domain.Dinosaur{
		Name:        dto.Name,
		Description: dto.Description,
	}
	return a.dinoUsecase.CreateDinosaur(newDinosaur)
}
