<p align="center">
    <img src="assets/moviechase.png" alt="Moviechase Logo" width="500" height="500">
</p>

# Moviechase API

This API serves as the backend for user authentication and management, as well as handling user-specific lists like favourites and watchlists. The database used is MySQL, and the API supports JSON Web Tokens (JWT) for secure user sessions.

## Features

- **User Authentication**: Register, login, and manage user accounts securely ( edit & delete).
- **JWT Integration**: Secure endpoints with token-based authentication.
- **Favourites Management**: Add, delete, and retrieve user favourites.
- **Watchlist Management**: Add, delete, and retrieve user watchlists.
- **Cloudinary Integration**: Upload and store user profile images in the cloud using Cloudinary.

## Prerequisites

- **Go**: Version 1.19+
- **MySQL**: Ensure the MySQL server is running and properly configured.
- **Environment Variables**: Configure a `.env` file or set the following variables:
  - `DB_HOST` - MySQL database host.
  - `DB_PORT` - MySQL database port.
  - `DB_USER` - MySQL username.
  - `DB_PASSWORD` - MySQL password.
  - `DB_NAME` - MySQL database name.
  - `JWT_SECRET` - Secret key for signing JWTs.

## Database Schema

### Users Table

- `id`
- `firstName`
- `lastName`
- `email`
- `password`
- `createdAt`
- `image`

### Favourites Table

- `userId`
- `id`
- `type`

### Watchlist Table

- `userId`
- `id`
- `type`

## Security

- **Password Hashing**: User passwords are hashed using a secure algorithm.
- **JWT**: Access to protected endpoints requires a valid JWT in the `Authorization` header.

## License

This project is licensed under the MIT License.
