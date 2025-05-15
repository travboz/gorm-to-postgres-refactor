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
		router.Post("/quests", CreateQuest(env))
		router.Get("/quests", GetAllQuests(env))
		// router.Get("/quests/{id}", GetQuestById(env))
		// router.Put("/quests/{id}", UpdateQuest(env))
		// router.Delete("/quests/{id}", DeleteQuestById(env))
	})

	return router
}
