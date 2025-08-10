# Event Management API

A simple REST API built using **Go (Golang)**, **Gin** framework, and **JWT authentication** for managing events and users.

## Features

* User registration & login
* JWT-based authentication
* Create and manage events
* Token verification middleware

## Tech Stack

* **Go** (Golang)
* **Gin** HTTP Web Framework
* **JWT** for authentication

## Project Structure

```
.
├── main.go
├── controllers
│   ├── userController.go
│   ├── eventController.go
├── models
│   ├── users.go
│   ├── event.go
├── utils
│   ├── token.go
└── README.md
```

## Setup & Installation

1. **Clone the repository:**

```bash
git clone https://github.com/yourusername/event-management-api.git
cd event-management-api
```

2. **Install dependencies:**

```bash
go mod tidy
```

3. **Configure Environment Variables:**

    * Set your database connection string.
    * Set your `JWT_SECRET_KEY`.
4. **Run the server:**

```bash
go run main.go
```

5. **API runs on:**

```
http://localhost:8081
```

## Authentication Flow

1. User logs in using `/login` endpoint.
2. Server returns a **JWT token**.
3. Client sends this token in the **Authorization header** as `Bearer <token>` for protected routes.

## Endpoints

### User

**POST /login**

* Request:

```json
{
  "EMAIL": "user@example.com",
  "PASSWORD": "password123"
}
```

* Response:

```json
{
  "token": "<jwt_token>"
}
```

### Events

**POST /events** *(Protected)*

* Headers:

```
Authorization: Bearer <jwt_token>
```

* Request:

```json
{
  "title": "My Event",
  "description": "An awesome event",
  "date": "2025-08-11"
}
```

* Response:

```json
{
  "message": "Event created",
  "event": {
    "title": "My Event",
    "description": "An awesome event",
    "date": "2025-08-11"
  }
}
```

## Token Utility Functions

* `GenerateToken(email string, id int64)`: Creates a JWT token.
* `VerifyToken(tokenString string)`: Verifies the token's validity.

## Notes

* Replace `secretKey` in `utils/token.go` with a secure value.
* Ensure your database is configured before running.
* Tokens expire after 2 hours (configurable).

## Contributing

Contributions are welcome! To contribute:

1. Fork this repository.
2. Create a new branch for your feature or bug fix.
3. Commit and push your changes.
4. Open a pull request describing your changes.

Please follow coding standards and include tests where applicable.

## License

MIT License
