# Event Management API

A simple and secure REST API built using **Go (Golang)**, the **Gin** framework, and **JWT authentication** for managing events and users.

## Features

* User registration and login
* JWT-based authentication with token expiry
* Event creation, listing, and management
* Middleware for token verification and route protection
* Modular code structure for scalability

## Tech Stack

* **Go** (Golang)
* **Gin** HTTP Web Framework
* **JWT** for authentication
* **MySQL** (or compatible) database

## Project Structure

```
.
├── main.go              # Application entry point
├── api-test             # Postman/test collections
├── db                   # Database connection and setup
├── middlewares          # Authentication and other middlewares
├── models               # Database models
├── routes               # API route definitions
├── utils                # Helper utilities (JWT, password hashing)
└── README.md
```

## Setup & Installation

1. **Clone the repository:**

```bash
git clone https://github.com/Rohitth10e/event-management-api.git
cd event-management-api
```

2. **Install dependencies:**

```bash
go mod tidy
```

3. **Configure environment variables:**

   * Set your database connection string in `db/connection.go`.
   * Update `JWT_SECRET_KEY` in `.env` or `utils/token.go`.
4. **Run the server:**

```bash
go run main.go
```

5. API available at:

```
http://localhost:8081
```

## Authentication Flow

1. **Signup:** POST `/signup` with email and password.
2. **Login:** POST `/login` to receive JWT token.
3. **Use Token:** Include `Authorization: Bearer <token>` for protected endpoints.
4. Token expires in **2 hours** (configurable in `utils/token.go`).

## Endpoints

### User

**POST /signup**

```json
{
  "EMAIL": "user@example.com",
  "PASSWORD": "password123"
}
```

**POST /login**

```json
{
  "EMAIL": "user@example.com",
  "PASSWORD": "password123"
}
```

*Response:* `{ "token": "<jwt_token>" }`

### Events (Protected)

**POST /events**
Headers: `Authorization: Bearer <jwt_token>`

```json
{
  "NAME": "My Event",
  "DESCRIPTION": "An awesome event",
  "LOCATION": "Bangalore"
}
```

## Notes for Frontend Developers

* Always send the token in the `Authorization` header for protected routes.
* Token expiry: 2 hours.
* Date fields should follow ISO 8601 (`YYYY-MM-DD` or `YYYY-MM-DDTHH:mm:ssZ`).
* For local testing, configure CORS in the backend if needed.
* Example workflow: User signs up → logs in → stores JWT in secure storage (not localStorage) → includes it in all API requests.

## Contributing

We welcome contributions! To contribute:

1. **Fork** the repository.
2. Create a feature branch: `git checkout -b feature-name`.
3. Commit changes with a descriptive message: `git commit -m "Description of changes"`.
4. Push your branch: `git push origin feature-name`.
5. Open a **Pull Request** explaining your changes.

### Contribution Guidelines

* Write clear, concise commit messages.
* Follow Go best practices and code formatting (`go fmt`).
* Update documentation where necessary.
* Add or update tests for any changes.

## Issues

* Report bugs or request features via the GitHub **Issues** tab.
* Include:

  * Steps to reproduce
  * Expected vs actual behavior
  * Environment details (OS, Go version, etc.)

## License

MIT License
