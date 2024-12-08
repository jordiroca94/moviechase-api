package favourites

import (
	"github.com/gorilla/mux"
)

func RegisterFavouritesRoutes(router *mux.Router, handler *FavouritesHandler) {
	router.HandleFunc("/favourites/add", handler.handleAddFavourite).Methods("POST")
}
