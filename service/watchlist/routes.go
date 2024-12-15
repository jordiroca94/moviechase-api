package watchlist

import (
	"github.com/gorilla/mux"
)

func RegisterWatchlistRoutes(router *mux.Router, handler *WatchlistHandler) {
	router.HandleFunc("/watchlist/add", handler.handleAddWatched).Methods("POST")
	router.HandleFunc("/watchlist/delete", handler.handleDeleteWatched).Methods("DELETE")
	router.HandleFunc("/watchlist", handler.handleGetWatchlist).Methods("GET")
	router.HandleFunc("/watched", handler.handleGetWatched).Methods("GET")
}
