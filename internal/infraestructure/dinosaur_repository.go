package infrastructure

import (
	"dinosaur-management/internal/domain"
	"sync"
)

type DinosaurRepositoryImpl struct {
	dinosaurs map[int]domain.Dinosaur
	mu        sync.RWMutex
}

func NewDinosaurRepository() *DinosaurRepositoryImpl {
	return &DinosaurRepositoryImpl{
		dinosaurs: make(map[int]domain.Dinosaur),
	}
}

func (r *DinosaurRepositoryImpl) ListDinosaurs() ([]domain.Dinosaur, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	dinos := make([]domain.Dinosaur, 0, len(r.dinosaurs))
	for _, dino := range r.dinosaurs {
		dinos = append(dinos, dino)
	}
	return dinos, nil
}

func (r *DinosaurRepositoryImpl) CreateDinosaur(dino domain.Dinosaur) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Simulate auto-incrementing ID
	dino.ID = len(r.dinosaurs) + 1
	r.dinosaurs[dino.ID] = dino
	return nil
}
