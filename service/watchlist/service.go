package watchlist

import "github.com/jordiroca94/moviechase-api/types"

type WatchlistService struct {
	repository *WatchlistRespository
}

func NewWatchlistService(repository *WatchlistRespository) *WatchlistService {
	return &WatchlistService{
		repository: repository,
	}
}

func (s *WatchlistService) AddWatched(watched types.FavouritesPayload) error {
	err := s.repository.AddWatched(watched)
	if err != nil {
		return err
	}
	return nil
}

func (s *WatchlistService) GetWatchedByMovieIDAndUserID(payload types.FavouritesPayload) error {
	_, err := s.repository.GetWatchedByMovieIDAndUserID(payload)
	if err != nil {
		return err
	}
	return nil
}

func (s *WatchlistService) DeleteWatched(watched types.FavouritesPayload) error {
	err := s.repository.DeleteWatched(watched)
	if err != nil {
		return err
	}
	return nil
}

func (s *WatchlistService) GetWatchlistByUserID(userID int) ([]types.FavouritesPayload, error) {
	watchlist, err := s.repository.GetWatchlistByUserID(userID)
	if err != nil {
		return nil, err
	}
	return watchlist, nil
}

func (s *WatchlistService) GetWatched(id int, userId int, favType types.FavouritesType) (*types.FavouritesPayload, error) {

	watched := types.FavouritesPayload{
		ID:     id,
		UserID: userId,
		Type:   favType,
	}
	result, err := s.repository.GetWatchedByMovieIDAndUserID(watched)
	if err != nil {
		return nil, err
	}
	return result, nil
}
