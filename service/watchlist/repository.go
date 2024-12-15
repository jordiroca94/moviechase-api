package watchlist

import (
	"database/sql"

	"github.com/jordiroca94/moviechase-api/types"
)

type WatchlistRespository struct {
	db *sql.DB
}

func NewWatchlistRepository(db *sql.DB) *WatchlistRespository {
	return &WatchlistRespository{db: db}
}

func (r *WatchlistRespository) AddWatched(watched types.FavouritesPayload) error {
	_, err := r.db.Exec("INSERT INTO watchlist (userId, id, type) VALUES (?, ?,?)", watched.UserID, watched.ID, watched.Type)
	if err != nil {
		return err
	}
	return nil
}

func (r *WatchlistRespository) GetWatchedByMovieIDAndUserID(watched types.FavouritesPayload) (*types.FavouritesPayload, error) {
	var result types.FavouritesPayload
	err := r.db.QueryRow("SELECT userId, id, type FROM watchlist WHERE id = ? AND userId = ? AND type = ?",
		watched.ID, watched.UserID, watched.Type).
		Scan(&result.UserID, &result.ID, &result.Type)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *WatchlistRespository) DeleteWatched(watched types.FavouritesPayload) error {
	_, err := r.db.Exec("DELETE FROM watchlist WHERE id = ? AND userId = ? AND type = ?", watched.ID, watched.UserID, watched.Type)
	if err != nil {
		return err
	}
	return nil
}

func (r *WatchlistRespository) GetWatchlistByUserID(userID int, typeFav string) ([]types.FavouritesPayload, error) {
	rows, err := r.db.Query("SELECT * FROM watchlist WHERE userId = ? AND type = ?", userID, typeFav)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var watchlist []types.FavouritesPayload
	for rows.Next() {
		var fav types.FavouritesPayload
		err := rows.Scan(&fav.UserID, &fav.ID, &fav.Type)
		if err != nil {
			return nil, err
		}
		watchlist = append(watchlist, fav)
	}
	return watchlist, nil
}
