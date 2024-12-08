package types

import "time"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

type FavouritesPayload struct {
	ID     int            `json:"id" validate:"required"`
	UserID int            `json:"user_id" validate:"required"`
	Type   FavouritesType `json:"type" validate:"required"`
}

type FavouritesType string

const (
	MoviesType FavouritesType = "movies"
	ShowsType  FavouritesType = "shows"
	PeopleType FavouritesType = "people"
)
