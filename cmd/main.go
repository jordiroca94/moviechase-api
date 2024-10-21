package main

//RESTful API  for managing user registration, login, and authentication with golang,echo.
// 1. SET UP  Golang with Echo  for routing and handling HTTP requests.
// 2.1 DATABASE: Use PostgreSQL or MySQL for storing user data.
// 2.2 Install GORM to simplify database interactions.
// 2.3 Create a User model to store fields like id, username, password_hash, email, and timestamps.
// 3. PASSWORD HASHING : Use a library like bcrypt for securely hashing
// 4.JWT Authentication : Use JWT for user authentication and authorization.
// 5 AUTHENTIFICATION MIDDLEWARE: Protect certain routes (like /profile) by using middleware to verify the JWT token.
// 6 ROUTES: Create routes for user registration, login, and profile.
// 6.1 POST /register: Accepts a request body with username, email, and password, hashes the password, and saves the user to the database.
// 6.2 POST /login: Checks if the user exists and if the password matches, then returns a JWT token.
// 6.3 GET /profile: A protected route that returns the logged-in user's details based on the JWT.

import (
	"github.com/jordiroca94/user-auth-api/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		panic(err)
	}

}
