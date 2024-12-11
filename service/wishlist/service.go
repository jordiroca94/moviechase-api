package wishlist

import "github.com/jordiroca94/moviechase-api/types"

type WishlistService struct {
	repository *WishlistRespository
}

func NewWishlistService(repository *WishlistRespository) *WishlistService {
	return &WishlistService{
		repository: repository,
	}
}

func (s *WishlistService) AddWished(wished types.FavouritesPayload) error {
	err := s.repository.AddWished(wished)
	if err != nil {
		return err
	}
	return nil
}

func (s *WishlistService) GetWishedByMovieIDAndUserID(payload types.FavouritesPayload) error {
	_, err := s.repository.GetWishedByMovieIDAndUserID(payload)
	if err != nil {
		return err
	}
	return nil
}

func (s *WishlistService) DeleteWished(wished types.FavouritesPayload) error {
	err := s.repository.DeleteWished(wished)
	if err != nil {
		return err
	}
	return nil
}

func (s *WishlistService) GetWishlistByUserID(userID int, typeFav string) ([]types.FavouritesPayload, error) {
	wishlist, err := s.repository.GetWishlistByUserID(userID, typeFav)
	if err != nil {
		return nil, err
	}
	return wishlist, nil
}

func (s *WishlistService) GetWished(id int, userId int, favType types.FavouritesType) (*types.FavouritesPayload, error) {

	wished := types.FavouritesPayload{
		ID:     id,
		UserID: userId,
		Type:   favType,
	}
	result, err := s.repository.GetWishedByMovieIDAndUserID(wished)
	if err != nil {
		return nil, err
	}
	return result, nil
}
