package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Group(func(r chi.Router) {
		router.Get("/quests", GetAllQuests)
		router.Get("/quest/{id}", GetQuestById)
		router.Post("/quest", CreateQuest)
		router.Put("/quest/{id}", UpdateQuest)
		router.Delete("/quest/{id}", DeleteQuestById)
	})

	return router
}
