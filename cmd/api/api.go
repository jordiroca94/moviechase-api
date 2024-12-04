package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jordiroca94/moviechase-api/service/user"
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
	userHandler := user.NewHandler(userService, userRepository)
	userHandler.RegisterRoutes(subrouter)

	fmt.Println("Server is running on port", s.addr)

	return http.ListenAndServe(s.addr, enableCors(router))
}
