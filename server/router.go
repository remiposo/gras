package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/remiposo/gras/infra/repository"
	"github.com/remiposo/gras/presentation"
	"github.com/remiposo/gras/usecase"
)

func NewRouter(db *sqlx.DB) http.Handler {
	r := chi.NewRouter()

	// set general middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", presentation.Health)

	taskRepo := repository.NewTaskRepo(db)
	taskUc := usecase.NewTaskUc(taskRepo)
	taskHandler := presentation.NewTaskHandler(taskUc)
	r.Post("/tasks", taskHandler.Add)
	r.Get("/tasks", taskHandler.List)

	userRepo := repository.NewUserRepo(db)
	userUc := usecase.NewUserUc(userRepo)
	userHandler := presentation.NewUserHandler(userUc)
	r.Post("/register", userHandler.Register)
	return r
}
