package api

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/jordiroca94/moviechase-api/service/favourites"
	"github.com/jordiroca94/moviechase-api/service/user"
	"github.com/jordiroca94/moviechase-api/service/watchlist"
)

func RegisterRoutes(router *mux.Router, db *sql.DB) {
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewHandler(userService)
	user.RegisterUserRoutes(subrouter, userHandler)

	favouritesRepository := favourites.NewFavouritesRepository(db)
	favouritesService := favourites.NewFavouritesService(favouritesRepository)
	favouritesHandler := favourites.NewHandler(favouritesService)
	favourites.RegisterFavouritesRoutes(subrouter, favouritesHandler)

	watchlistRepository := watchlist.NewWatchlistRepository(db)
	watchlistService := watchlist.NewWatchlistService(watchlistRepository)
	watchlistHandler := watchlist.NewHandler(watchlistService)
	watchlist.RegisterWatchlistRoutes(subrouter, watchlistHandler)

}
