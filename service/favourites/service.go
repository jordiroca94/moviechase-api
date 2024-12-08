package favourites

import "github.com/jordiroca94/moviechase-api/types"

type FavouritesService struct {
	repository *FavouritesRespository
}

func NewFavouritesService(repository *FavouritesRespository) *FavouritesService {
	return &FavouritesService{
		repository: repository,
	}
}

func (s *FavouritesService) AddFavourite(favourite types.FavouritesPayload) error {
	err := s.repository.AddFavourite(favourite)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavouritesService) GetFavouriteByMovieIDAndUserID(payload types.FavouritesPayload) error {
	err := s.repository.GetFavouriteByMovieIDAndUserID(payload)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavouritesService) DeleteFavourite(favourite types.FavouritesPayload) error {
	err := s.repository.DeleteFavourite(favourite)
	if err != nil {
		return err
	}
	return nil
}

func (s *FavouritesService) GetFavouritesByUserID(userID int, typeFav string) ([]types.FavouritesPayload, error) {
	favourites, err := s.repository.GetFavouritesByUserID(userID, typeFav)
	if err != nil {
		return nil, err
	}
	return favourites, nil
}
