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

func (r *FavouritesRespository) GetFavouriteByMovieIDAndUserID(favourite types.FavouritesPayload) (*types.FavouritesPayload, error) {
	var result types.FavouritesPayload
	err := r.db.QueryRow("SELECT userId, id, type FROM favourites WHERE id = ? AND userId = ? AND type = ?",
		favourite.ID, favourite.UserID, favourite.Type).
		Scan(&result.UserID, &result.ID, &result.Type)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *FavouritesRespository) DeleteFavourite(favourite types.FavouritesPayload) error {
	_, err := r.db.Exec("DELETE FROM favourites WHERE id = ? AND userId = ? AND type = ?", favourite.ID, favourite.UserID, favourite.Type)
	if err != nil {
		return err
	}
	return nil
}

// GetFavouritesByUserID with type as a parameter to get all favourites of a user
func (r *FavouritesRespository) GetFavouritesByUserID(userID int, typeFav string) ([]types.FavouritesPayload, error) {
	rows, err := r.db.Query("SELECT * FROM favourites WHERE userId = ? AND type = ?", userID, typeFav)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var favourites []types.FavouritesPayload
	for rows.Next() {
		var fav types.FavouritesPayload
		err := rows.Scan(&fav.UserID, &fav.ID, &fav.Type)
		if err != nil {
			return nil, err
		}
		favourites = append(favourites, fav)
	}
	return favourites, nil
}
