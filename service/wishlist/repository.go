package wishlist

import (
	"database/sql"

	"github.com/jordiroca94/moviechase-api/types"
)

type WishlistRespository struct {
	db *sql.DB
}

func NewWishlistRepository(db *sql.DB) *WishlistRespository {
	return &WishlistRespository{db: db}
}

func (r *WishlistRespository) AddWished(wished types.FavouritesPayload) error {
	_, err := r.db.Exec("INSERT INTO wishlist (userId, id, type) VALUES (?, ?,?)", wished.UserID, wished.ID, wished.Type)
	if err != nil {
		return err
	}
	return nil
}

func (r *WishlistRespository) GetWishedByMovieIDAndUserID(wished types.FavouritesPayload) (*types.FavouritesPayload, error) {
	var result types.FavouritesPayload
	err := r.db.QueryRow("SELECT userId, id, type FROM wishlist WHERE id = ? AND userId = ? AND type = ?",
		wished.ID, wished.UserID, wished.Type).
		Scan(&result.UserID, &result.ID, &result.Type)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *WishlistRespository) DeleteWished(wished types.FavouritesPayload) error {
	_, err := r.db.Exec("DELETE FROM wishlist WHERE id = ? AND userId = ? AND type = ?", wished.ID, wished.UserID, wished.Type)
	if err != nil {
		return err
	}
	return nil
}

func (r *WishlistRespository) GetWishlistByUserID(userID int, typeFav string) ([]types.FavouritesPayload, error) {
	rows, err := r.db.Query("SELECT * FROM wishlist WHERE userId = ? AND type = ?", userID, typeFav)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var wishlist []types.FavouritesPayload
	for rows.Next() {
		var fav types.FavouritesPayload
		err := rows.Scan(&fav.UserID, &fav.ID, &fav.Type)
		if err != nil {
			return nil, err
		}
		wishlist = append(wishlist, fav)
	}
	return wishlist, nil
}
