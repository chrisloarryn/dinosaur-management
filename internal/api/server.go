package api

import (
	"dinosaur-management/internal/adapter"
	"dinosaur-management/internal/domain"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	e           *echo.Echo
	dinoAdapter adapter.DinosaurAdapter
}

func NewServer(dinoAdapter adapter.DinosaurAdapter) *Server {
	e := echo.New()
	e.Validator = &domain.CustomValidator{Validator: validator.New()}
	return &Server{e: e, dinoAdapter: dinoAdapter}
}

func (s *Server) Start(address string) error {
	s.e.GET("/dinosaurs", s.listDinosaurs)
	s.e.POST("/dinosaurs", s.createDinosaur)

	return s.e.Start(address)
}

func (s *Server) listDinosaurs(c echo.Context) error {
	dinosaurs, err := s.dinoAdapter.ListDinosaurs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	return c.JSON(http.StatusOK, dinosaurs)
}

func (s *Server) createDinosaur(c echo.Context) error {
	var newDino adapter.DinosaurDTO

	// validate not empty body and bind to newDino
	if err := c.Bind(&newDino); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// validate newDino
	if err := c.Validate(newDino); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := s.dinoAdapter.CreateDinosaur(newDino)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create dinosaur"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "Dinosaur created successfully"})
}
