package favourites

import (
	"github.com/gorilla/mux"
)

func RegisterFavouritesRoutes(router *mux.Router, handler *FavouritesHandler) {
	router.HandleFunc("/favourites/add", handler.handleAddFavourite).Methods("POST")
	router.HandleFunc("/favourites/delete", handler.handleDeleteFavourite).Methods("DELETE")
	router.HandleFunc("/favourites", handler.handleGetFavourites).Methods("GET")
	router.HandleFunc("/favourite", handler.handleGetFavourite).Methods("GET")
}
