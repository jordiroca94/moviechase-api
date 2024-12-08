package favourites

import (
	"database/sql"

	"github.com/jordiroca94/moviechase-api/types"
)

type FavouritesRespository struct {
	db *sql.DB
}

func NewFavouritesRepository(db *sql.DB) *FavouritesRespository {
	return &FavouritesRespository{db: db}
}

func (r *FavouritesRespository) AddFavourite(favourite types.FavouritesPayload) error {
	_, err := r.db.Exec("INSERT INTO favourites (userId, id, type) VALUES (?, ?,?)", favourite.UserID, favourite.ID, favourite.Type)
	if err != nil {
		return err
	}
	return nil
}

func (r *FavouritesRespository) GetFavouriteByMovieIDAndUserID(favourite types.FavouritesPayload) error {
	err := r.db.QueryRow("SELECT * FROM favourites WHERE id = ? AND userId = ? AND type = ?", favourite.ID, favourite.UserID, favourite.Type).Scan(&favourite.UserID, &favourite.ID, &favourite.Type)
	if err != nil {
		return err
	}
	return nil
}
