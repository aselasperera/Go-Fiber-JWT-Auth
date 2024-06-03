# Go Fiber JWT Authentication

This project demonstrates a basic user authentication and authorization system using Go Fiber and JWT (JSON Web Token). Users can log in with a username and password to receive a JWT token, which can be used to access protected routes.

## Features

- User login and JWT token generation
- Protected route that requires a valid JWT token
- Basic username and password authentication

## Project Structure

├── main.go
├── go.mod
├── go.sum
├── handler
│ ├── auth.go
├── middleware
│ ├── jwt.go
└── README.md


## Getting Started

### Prerequisites

- Go (1.16 or higher)
- Git

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/aselasperera/Go-Fiber-JWT-Auth
    cd your-repo-name
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Run the application:

    ```bash
    go run main.go
    ```

The server will start on `localhost:4000`.

## API Endpoints

### 1. Login

- **URL**: `/login`
- **Method**: `POST`
- **Headers**: `Content-Type: application/json`
- **Body**:

    ```json
    {
        "username": "Chek",
        "password": "123456"
    }
    ```

- **Success Response**:

    - **Code**: 200 OK
    - **Content**: JWT token string

- **Error Response**:

    - **Code**: 401 Unauthorized
    - **Content**: `Invalid credentials`

### 2. Protected Route

- **URL**: `/protected`
- **Method**: `GET`
- **Headers**: `Authorization: Bearer <your-jwt-token>`

- **Success Response**:

    - **Code**: 200 OK
    - **Content**: `Welcome to the protected area`

- **Error Response**:

    - **Code**: 401 Unauthorized
    - **Content**: `Missing authorization header` or `Invalid token`

## Testing the Application

### Using Postman

1. **Login Request**:

    - **URL**: `http://localhost:4000/login`
    - **Method**: `POST`
    - **Body**:

        ```json
        {
            "username": "Chek",
            "password": "123456"
        }
        ```

2. **Protected Request**:

    - **URL**: `http://localhost:4000/protected`
    - **Method**: `GET`
    - **Headers**: `Authorization: Bearer <your-jwt-token>`

### Using `curl`

1. **Login Request**:

    ```bash
    curl -X POST http://localhost:4000/login -H "Content-Type: application/json" -d '{"username": "Chek", "password": "123456"}'
    ```

    This will return a JWT token if the credentials are correct.

2. **Protected Request**:

    Replace `<your-jwt-token>` with the token received from the login request:

    ```bash
    curl -X GET http://localhost:4000/protected -H "Authorization: Bearer <your-jwt-token>"
    ```

    If the token is valid, you should receive the message "Welcome to the protected area".

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Fiber](https://gofiber.io/) - Express inspired web framework written in Go
- [golang-jwt](https://github.com/golang-jwt/jwt) - A Go implementation of JSON Web Tokens
