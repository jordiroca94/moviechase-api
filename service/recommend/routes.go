package recommend

import (
	"github.com/gorilla/mux"
)

func RegisterRecommendRoutes(router *mux.Router, handler *RecommendHandler) {
	router.HandleFunc("/recommend", handler.handleGetRecommendation).Methods("POST")
}
