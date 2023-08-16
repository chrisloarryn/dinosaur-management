package main

import (
	"dinosaur-management/internal/adapter"
	"dinosaur-management/internal/api"
	infrastructure "dinosaur-management/internal/infraestructure"
	"dinosaur-management/internal/usecase"
)

func main() {
	dinoAdapter := adapter.NewDinosaurAdapter(usecase.NewDinosaurUsecase(infrastructure.NewDinosaurRepository()))
	server := api.NewServer(dinoAdapter)
	err := server.Start(":8080")
	if err != nil {
		return
	}
}
