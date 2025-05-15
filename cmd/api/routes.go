package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/travboz/go-quest/internal/env"
)

func routes(env *env.Env) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Group(func(r chi.Router) {
		// router.Get("/quests", GetAllQuests(env))
		// router.Get("/quest/{id}", GetQuestById(env))
		router.Post("/quest", CreateQuest(env))
		// router.Put("/quest/{id}", UpdateQuest(env))
		// router.Delete("/quest/{id}", DeleteQuestById(env))
	})

	return router
}
