package watchlist

import (
	"github.com/gorilla/mux"
)

func RegisterWatchlistRoutes(router *mux.Router, handler *WatchlistHandler) {
	router.HandleFunc("/watchlist/add", handler.handleAddWatched).Methods("POST")
	router.HandleFunc("/watchlist/delete", handler.handleDeleteWatched).Methods("DELETE")
	router.HandleFunc("/watchlist", handler.handleGetWatchlist).Queries("user_id", "{user_id}", "type", "{type}").Methods("GET")
	router.HandleFunc("/watched", handler.handleGetWatched).Queries("user_id", "{user_id}", "type", "{type}", "id", "{id}").Methods("GET")
}
