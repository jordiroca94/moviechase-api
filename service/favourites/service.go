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
