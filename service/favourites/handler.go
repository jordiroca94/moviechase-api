package favourites

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/jordiroca94/moviechase-api/types"
	"github.com/jordiroca94/moviechase-api/utils"
)

type FavouritesHandler struct {
	service *FavouritesService
}

func NewHandler(service *FavouritesService) *FavouritesHandler {
	return &FavouritesHandler{
		service: service,
	}
}

func (h *FavouritesHandler) handleAddFavourite(w http.ResponseWriter, r *http.Request) {
	var payload types.FavouritesPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", errors))
		return
	}

	if payload.Type != types.MoviesType && payload.Type != types.ShowsType && payload.Type != types.PeopleType {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("%s is an invalid type", payload.Type))
		return
	}

	err := h.service.GetFavouriteByMovieIDAndUserID(payload)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("%s with id %s is already in favourites", payload.Type, strconv.Itoa(payload.ID)))
		return
	}

	err = h.service.AddFavourite(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "added to favourites successfully"})
}

func (h *FavouritesHandler) handleDeleteFavourite(w http.ResponseWriter, r *http.Request) {
	var payload types.FavouritesPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", errors))
		return
	}

	if payload.Type != types.MoviesType && payload.Type != types.ShowsType && payload.Type != types.PeopleType {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("%s is an invalid type", payload.Type))
		return
	}

	err := h.service.GetFavouriteByMovieIDAndUserID(payload)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("%s with id %s is not in favourites", payload.Type, strconv.Itoa(payload.ID)))
		return
	}

	err = h.service.DeleteFavourite(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "removed from favourites successfully"})
}
