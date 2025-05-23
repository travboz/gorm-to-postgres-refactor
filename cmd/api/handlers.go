package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/travboz/go-quest/internal/env"
	"github.com/travboz/go-quest/internal/models"
	"github.com/travboz/go-quest/internal/utils"
)

func GetAllQuests(env *env.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(utils.ContentType, utils.ContentJSON)

		quests, err := models.GetAllQuests(env.DB)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = utils.RespondWithJSON(w, http.StatusOK, quests)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func GetQuestById(env *env.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "unable to parse id")
			return
		}

		quest, err := models.GetQuestByID(env.DB, id)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrRecordNotFound):
				utils.RespondWithError(w, http.StatusNotFound, "quest not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}

			return
		}

		if err = utils.RespondWithJSON(w, http.StatusOK, quest); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}
}

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

		if err = utils.RespondWithJSON(w, http.StatusCreated, quest); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}
}

type UpdateQuestPayload struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Reward      int    `json:"reward" validate:"required"`
}

func UpdateQuest(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "unable to parse id")
			return
		}

		// Getting quest to update
		quest, err := models.GetQuestByID(env.DB, id)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrRecordNotFound):
				utils.RespondWithError(w, http.StatusNotFound, "quest not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}

			return
		}

		// Reading the JSON update payload
		var input UpdateQuestPayload

		body, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(body, &input)

		validate = validator.New()
		if err := validate.Struct(input); err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "validation error")
			return
		}

		// Updating values
		quest.Title = input.Title
		quest.Description = input.Description
		quest.Reward = input.Reward

		// Saving the update in our db
		err = models.UpdateQuest(env.DB, quest)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrEditConflict):
				utils.RespondWithError(w, http.StatusConflict, models.ErrEditConflict.Error())
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		if err = utils.RespondWithJSON(w, http.StatusAccepted, quest); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}
}

func DeleteQuestById(env *env.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "unable to parse id")
			return
		}

		err = models.DeleteQuest(env.DB, id)
		if err != nil {
			switch {
			case errors.Is(err, models.ErrRecordNotFound):
				utils.RespondWithError(w, http.StatusNotFound, "quest not found")
			default:
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			}

			return
		}

		if err = utils.RespondWithJSON(w, http.StatusAccepted, map[string]string{"message": "quest deleted successfully"}); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		}
	}
}
