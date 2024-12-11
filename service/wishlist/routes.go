package wishlist

import (
	"github.com/gorilla/mux"
)

func RegisterWishlistRoutes(router *mux.Router, handler *WishlistHandler) {
	router.HandleFunc("/wishlist/add", handler.handleAddWished).Methods("POST")
	router.HandleFunc("/wishlist/delete", handler.handleDeleteWished).Methods("DELETE")
	router.HandleFunc("/wishlist", handler.handleGetWishlist).Methods("GET")
	router.HandleFunc("/wished", handler.handleGetWished).Methods("GET")
}
