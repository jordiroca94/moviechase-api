package user

import (
	"database/sql"
	"fmt"

	"github.com/jordiroca94/moviechase-api/types"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByEmail(email string) (*types.User, error) {
	rows, err := r.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt, &user.Image)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByID(id int) (*types.User, error) {

	rows, err := r.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (r *UserRepository) CreateUser(user *types.User) error {

	_, err := r.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(userID int, newUserData *types.UpdateUserPayload) error {
	_, err := r.db.Exec("UPDATE users SET firstName = ?, lastName = ?, email = ? WHERE id = ?", newUserData.FirstName, newUserData.LastName, newUserData.Email, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUserImage(id int, image string) error {
	_, err := r.db.Exec("UPDATE users SET image = ? WHERE id = ?", image, id)
	if err != nil {
		return err
	}
	return nil
}
