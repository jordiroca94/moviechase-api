package watchlist

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/jordiroca94/moviechase-api/types"
	"github.com/jordiroca94/moviechase-api/utils"
)

type WatchlistHandler struct {
	service *WatchlistService
}

func NewHandler(service *WatchlistService) *WatchlistHandler {
	return &WatchlistHandler{
		service: service,
	}
}

func (h *WatchlistHandler) handleAddWatched(w http.ResponseWriter, r *http.Request) {
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

	err := h.service.GetWatchedByMovieIDAndUserID(payload)
	if err == nil {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("%s with id %s is already in watchlist", payload.Type, strconv.Itoa(payload.ID)))
		return
	}

	err = h.service.AddWatched(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "added to watchlist successfully"})
}

func (h *WatchlistHandler) handleDeleteWatched(w http.ResponseWriter, r *http.Request) {
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

	err := h.service.GetWatchedByMovieIDAndUserID(payload)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("%s with id %s is not in watchlist", payload.Type, strconv.Itoa(payload.ID)))
		return
	}

	err = h.service.DeleteWatched(payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "removed from watchlist successfully"})
}

func (h *WatchlistHandler) handleGetWatchlist(w http.ResponseWriter, r *http.Request) {
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

	watchlist, err := h.service.GetWatchlistByUserID(userID, typeFav)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, watchlist)
}

func (h *WatchlistHandler) handleGetWatched(w http.ResponseWriter, r *http.Request) {
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

	watched, err := h.service.GetWatched(id, userID, types.FavouritesType(typeFav))
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("%s with id %s is not in watchlist", typeFav, strconv.Itoa(id)))
		return
	}
	utils.WriteJSON(w, http.StatusOK, watched)
}
