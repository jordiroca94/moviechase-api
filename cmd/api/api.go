package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jordiroca94/moviechase-api/service/favourites"
	"github.com/jordiroca94/moviechase-api/service/user"
	"github.com/jordiroca94/moviechase-api/service/wishlist"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Middleware function for handling CORS
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If it's a preflight request, we end it here
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userRepository := user.NewUserrepository(s.db)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewHandler(userService)
	user.RegisterUserRoutes(subrouter, userHandler)

	favouritesRepository := favourites.NewFavouritesRepository(s.db)
	favouritesService := favourites.NewFavouritesService(favouritesRepository)
	favouritesHandler := favourites.NewHandler(favouritesService)
	favourites.RegisterFavouritesRoutes(subrouter, favouritesHandler)

	wishlistRepository := wishlist.NewWishlistRepository(s.db)
	wishlistService := wishlist.NewWishlistService(wishlistRepository)
	wishlistHandler := wishlist.NewHandler(wishlistService)
	wishlist.RegisterWishlistRoutes(subrouter, wishlistHandler)

	router.HandleFunc("/routes", func(w http.ResponseWriter, r *http.Request) {
		var routes []map[string]string

		err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			path, err := route.GetPathTemplate()
			if err != nil {
				return nil
			}
			methods, err := route.GetMethods()
			if err != nil {
				methods = []string{"ALL"}
			}
			routes = append(routes, map[string]string{
				"path":    path,
				"methods": fmt.Sprintf("%v", methods),
			})
			return nil
		})
		if err != nil {
			http.Error(w, "Error listing routes", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		prettyJSON, err := json.MarshalIndent(routes, "", "  ")
		if err != nil {
			http.Error(w, "Error formatting JSON", http.StatusInternalServerError)
			return
		}
		w.Write(prettyJSON)
	})

	fmt.Println("Server is running on port", s.addr)

	return http.ListenAndServe(s.addr, enableCors(router))
}
