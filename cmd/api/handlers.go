package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/travboz/go-quest/internal/env"
	"github.com/travboz/go-quest/internal/models"
	"github.com/travboz/go-quest/internal/utils"
)

// func GetAllQuests(env *env.Env) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set(utils.ContentType, utils.ContentJSON)

// 		var quests []models.Quest
// 		models.DB.Find(&quests)

// 		json.NewEncoder(w).Encode(quests)
// 	}
// }

// func GetQuestById(env *env.Env) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set(utils.ContentType, utils.ContentJSON)

// 		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
// 		if err != nil {
// 			utils.RespondWithError(w, http.StatusInternalServerError, "unable to parse id")
// 			return
// 		}

// 		var quest models.Quest

// 		if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
// 			utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
// 			return
// 		}

// 		json.NewEncoder(w).Encode(quest)
// 	}
// }

var validate *validator.Validate

type CreateQuestPayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Reward      int    `json:"reward" validate:"required"`
}

func CreateQuest(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input CreateQuestPayload

		body, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(body, &input)

		validate = validator.New()
		if err := validate.Struct(input); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "validation error")
			return
		}

		quest := models.Quest{
			Title:       input.Title,
			Description: input.Description,
			Reward:      input.Reward,
		}

		err := models.CreateNewQuest(env.DB, quest)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set(utils.ContentType, utils.ContentJSON)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(quest)
	}
}

// type UpdateQuestPayload struct {
// 	Title       string `json:"title" validate:"required"`
// 	Description string `json:"description" validate:"required"`
// 	Reward      int    `json:"reward" validate:"required"`
// }

// func UpdateQuest(env *env.Env) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set(utils.ContentType, utils.ContentJSON)

// 		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
// 		if err != nil {
// 			utils.RespondWithError(w, http.StatusInternalServerError, "unable to parse id")
// 			return
// 		}

// 		// Getting quest to update
// 		var quest models.Quest

// 		if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
// 			utils.RespondWithError(w, http.StatusNotFound, "quest not found")
// 			return
// 		}

// 		// Reading the JSON update payload
// 		var input UpdateQuestPayload

// 		body, _ := io.ReadAll(r.Body)
// 		_ = json.Unmarshal(body, &input)

// 		validate = validator.New()
// 		if err := validate.Struct(input); err != nil {
// 			utils.RespondWithError(w, http.StatusBadRequest, "validation Error")
// 			return
// 		}

// 		// Updating values
// 		quest.Title = input.Title
// 		quest.Description = input.Description
// 		quest.Reward = input.Reward

// 		// Saving the update in our db
// 		models.DB.Save(&quest)

// 		json.NewEncoder(w).Encode(quest)
// 	}
// }

// func DeleteQuestById(env *env.Env) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set(utils.ContentType, utils.ContentJSON)

// 		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
// 		if err != nil {
// 			utils.RespondWithError(w, http.StatusInternalServerError, "unable to parse id")
// 			return
// 		}

// 		// Getting quest to update
// 		var quest models.Quest

// 		if err := models.DB.Where("id = ?", id).First(&quest).Error; err != nil {
// 			utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
// 			return
// 		}

// 		models.DB.Delete(&quest)

// 		w.WriteHeader(http.StatusNoContent)
// 		json.NewEncoder(w).Encode(quest)
// 	}
// }
