# Go Authentication Server

This is a simple authentication server built using **Go** and the **net/http** package. The server supports user registration, login, session management, and protected routes.

## Features
- User registration with hashed passwords (bcrypt)
- User login with session management
- Middleware for authentication
- SQLite database for storing users and sessions
- HTML templates for login and registration
- Static files (CSS, JS)

## Project Structure
```
go-auth-server/
│── main.go
│── handlers/
│   ├── auth.go
│   ├── home.go
│   ├── protected.go
│── middleware/
│   ├── auth.go
│   ├── csrf.go
│   ├── rate_limit.go
│   ├── session.go
│── templates/
│   ├── login.html
│   ├── register.html
│   ├── home.html
│── static/
│   ├── styles.css
│── database/
│   ├── db.go
│── go.mod
│── README.md
```

## Setup & Installation

### Prerequisites
- Install [Go](https://go.dev/doc/install)
- Install [SQLite3](https://www.sqlite.org/download.html)

### Clone the Repository
```sh
git clone https://github.com/yourusername/go-auth-server.git
cd go-auth-server
```

### Install Dependencies
```sh
go mod tidy
```

### Set Up Environment Variables
Create a `.env` file (optional but recommended):
```sh
echo "DATABASE_URL=./users.db" > .env
```
Or set it manually:
```sh
export DATABASE_URL=./users.db
```

### Run the Server
```sh
go run main.go
```

The server should be running at **`http://localhost:8080`**.

## API Endpoints
| Method | Endpoint       | Description        |
|--------|--------------|-------------------|
| `GET`  | `/`          | Home Page         |
| `GET`  | `/login`     | Login Page        |
| `POST` | `/login`     | Login User        |
| `GET`  | `/register`  | Register Page     |
| `POST` | `/register`  | Register User     |
| `GET`  | `/logout`    | Logout User       |
| `GET`  | `/dashboard` | Protected Route   |

## Database Schema
### Users Table
```sql
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);
```
### Sessions Table
```sql
CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    session_token TEXT UNIQUE NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## Deployment
### Deploy to Railway (Optional)
1. Push your code to GitHub
2. Connect the repository to [Railway](https://railway.app/)
3. Add `DATABASE_URL` as an environment variable
4. Deploy 🚀

## License
MIT License

## Author
Developed by **Anvarjon**

