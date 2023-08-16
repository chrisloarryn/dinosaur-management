package domain

// Dinosaur represents a dinosaur entity.
type Dinosaur struct {
	ID          int
	Name        string
	Description string
	// Add more fields as needed
}

// DinosaurUsecase is the interface that use cases in the domain layer implement.
type DinosaurUsecase interface {
	ListDinosaurs() ([]Dinosaur, error)
	CreateDinosaur(dino Dinosaur) error
}

// DinosaurRepository is the interface that repositories in the domain layer implement.
type DinosaurRepository interface {
	ListDinosaurs() ([]Dinosaur, error)
	CreateDinosaur(dino Dinosaur) error
}

// DinosaurService provides the business logic for dinosaur use cases.
type DinosaurService struct {
	repo DinosaurRepository
}

// NewDinosaurService creates a new instance of DinosaurService.
func NewDinosaurService(repo DinosaurRepository) *DinosaurService {
	return &DinosaurService{repo: repo}
}

// ListDinosaurs retrieves a list of all dinosaurs.
func (s *DinosaurService) ListDinosaurs() ([]Dinosaur, error) {
	dinosaurs, err := s.repo.ListDinosaurs()
	if err != nil {
		return nil, err
	}
	return dinosaurs, nil
}

// CreateDinosaur creates a new dinosaur.
func (s *DinosaurService) CreateDinosaur(dino Dinosaur) error {
	// You can perform validation or other business logic here before creating the dinosaur
	return s.repo.CreateDinosaur(dino)
}
