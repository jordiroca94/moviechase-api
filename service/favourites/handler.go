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

func (h *FavouritesHandler) handleGetFavourites(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user_id"))
		return
	}

	typeFav := r.URL.Query().Get("type")
	if types.FavouritesType(typeFav) != types.MoviesType && types.FavouritesType(typeFav) != types.ShowsType && types.FavouritesType(typeFav) != types.PeopleType {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("%s is an invalid type", typeFav))
		return
	}

	favourites, err := h.service.GetFavouritesByUserID(userID, typeFav)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, favourites)
}

func (h *FavouritesHandler) handleGetFavourite(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user_id"))
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid user_id"))
		return
	}

	typeFav := r.URL.Query().Get("type")
	if types.FavouritesType(typeFav) != types.MoviesType && types.FavouritesType(typeFav) != types.ShowsType && types.FavouritesType(typeFav) != types.PeopleType {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("%s is an invalid type", typeFav))
		return
	}

	favourite, err := h.service.GetFavourite(id, userID, types.FavouritesType(typeFav))
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("%s with id %s is not in favourites", typeFav, strconv.Itoa(id)))
		return
	}
	utils.WriteJSON(w, http.StatusOK, favourite)
}
