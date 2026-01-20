Movie_box 🎬
A robust backend movie management application built with Go, featuring JWT authentication, database migrations, and a RESTful API for movie data management.
Features

RESTful API: Well-structured endpoints for movie data operations
JWT Authentication: Secure token-based authentication system
Database Migrations: Managed schema changes using SQL migrations
CRUD Operations: Complete Create, Read, Update, Delete functionality for movies
User Management: User registration and authentication
Email Integration: Sending emails for user activation and notifications
Rate Limiting: API request rate limiting for security
Graceful Shutdown: Proper server shutdown handling
Structured Logging: Comprehensive error handling and logging
Cross-Origin Support: CORS configuration for frontend integration

Tech Stack

Language: Go 82.5%
Shell Scripts: 7.7%
Go Template: 5.3%
Makefile: 4.5%
Database: PostgreSQL (with migrations support)
Authentication: JWT (JSON Web Tokens)

Project Structure
movie_box_llife/
├── cmd/                    # Application entry points
├── internal/              # Private application code
├── migrations/            # Database migration files
├── remote/setup/          # Remote server setup scripts
├── vendor/                # Go dependencies
├── .gitignore
├── exit
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── learnam.pem           # SSL certificate
├── Makefile              # Build and deployment commands
└── .profile              # Shell profile configuration
Prerequisites

Go 1.21 or higher
PostgreSQL database
Make (for using Makefile commands)

Installation

Clone the repository:

bashgit clone https://github.com/OB-Peter/movie_box_llife.git
cd movie_box_llife

Install dependencies:

bashgo mod download

Set up environment variables:
Create a .env file or set the following environment variables:

bashexport DB_DSN="postgres://username:password@localhost/movie_box?sslmode=disable"
export JWT_SECRET="your-secret-key-here"
export SMTP_HOST="smtp.example.com"
export SMTP_PORT="587"
export SMTP_USERNAME="your-email@example.com"
export SMTP_PASSWORD="your-email-password"

Run database migrations:

bashmake db/migrations/up
# or manually
migrate -path=./migrations -database=$DB_DSN up

Build and run the application:

bashmake run
# or
go run ./cmd/api
Available Make Commands
bashmake help              # Show available commands
make run               # Run the application
make build             # Build the application binary
make db/migrations/up  # Run database migrations
make db/migrations/down # Rollback database migrations
make test              # Run tests
make deploy            # Deploy to production
API Endpoints
Authentication

POST /v1/users - Register a new user
POST /v1/tokens/authentication - Login and receive JWT token
POST /v1/tokens/activation - Activate user account

Movies

GET /v1/movies - List all movies (with filtering, sorting, pagination)
GET /v1/movies/:id - Get a specific movie
POST /v1/movies - Create a new movie (requires authentication)
PATCH /v1/movies/:id - Update a movie (requires authentication)
DELETE /v1/movies/:id - Delete a movie (requires authentication)

Health Check

GET /v1/healthcheck - Check API health status

Database Schema
The application uses SQL migrations to manage the database schema. Migration files are located in the migrations/ directory and include:

Movies table with fields for title, year, runtime, genres, etc.
Users table for authentication
Tokens table for JWT and activation tokens
Permissions table for role-based access control

Authentication
The API uses JWT (JSON Web Tokens) for authentication. To access protected endpoints:

Register a user account via POST /v1/users
Activate your account using the activation token sent to your email
Login via POST /v1/tokens/authentication to receive a JWT
Include the JWT in the Authorization header for subsequent requests:

   Authorization: Bearer <your-jwt-token>
Deployment
The application includes deployment scripts for remote servers:
bashmake deploy
This will build the application and deploy it to the configured remote server using the scripts in remote/setup/.
Configuration
Key configuration options can be set via environment variables or command-line flags:

PORT - Server port (default: 4000)
ENV - Environment (development/staging/production)
DB_DSN - PostgreSQL connection string
DB_MAX_OPEN_CONNS - Maximum open database connections
DB_MAX_IDLE_CONNS - Maximum idle database connections
RATE_LIMIT_ENABLED - Enable/disable rate limiting
RATE_LIMIT_RPS - Requests per second limit
SMTP_* - Email configuration

Development
Running Tests
bashgo test ./...
Code Formatting
bashgo fmt ./...
Linting
bashgolangci-lint run
Contributing
Contributions are welcome! Please follow these steps:

Fork the repository
Create a feature branch (git checkout -b feature/amazing-feature)
Commit your changes (git commit -m 'Add amazing feature')
Push to the branch (git push origin feature/amazing-feature)
Open a Pull Request

License
This project is licensed under the MIT License - see the LICENSE file for details.
Acknowledgments

Built following best practices from "Let's Go Further" by Alex Edwards
Uses the Gorilla toolkit for routing and middleware
PostgreSQL for reliable data storage

Contact
Oluyemi Boluwatife Peter

GitHub: @OB-Peter
Repository: movie_box_llife
