package favourites

import (
	"github.com/gorilla/mux"
)

func RegisterFavouritesRoutes(router *mux.Router, handler *FavouritesHandler) {
	router.HandleFunc("/favourites/add", handler.handleAddFavourite).Methods("POST")
	router.HandleFunc("/favourites/delete", handler.handleDeleteFavourite).Methods("DELETE")
	router.HandleFunc("/favourites", handler.handleGetFavourites).Queries("user_id", "{user_id}", "type", "{type}").Methods("GET")
	router.HandleFunc("/favourite", handler.handleGetFavourite).Queries("user_id", "{user_id}", "type", "{type}", "id", "{id}").Methods("GET")
}
