# Simple REST API with Go, GORM, Gin, and Aiven PostgreSQL

This project is a RESTful API built with Go, using the Gin framework for handling HTTP requests, GORM as the ORM for database interactions, and Aiven PostgreSQL for the database.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete resources.
- **Database Integration**: Seamless integration with Aiven PostgreSQL using GORM.
- **JSON Responses**: Serve and consume data in structured JSON format.
- **Error Handling**: Robust error handling for a smooth user experience.

## Technologies Used

- **Go**: Programming language used for developing the API.
- **Gin**: Web framework for handling HTTP requests.
- **GORM**: ORM library for database interactions.
- **Aiven PostgreSQL**: Managed PostgreSQL database hosted on Aiven.

## Getting Started

### Prerequisites

- Go installed on your machine.
- Aiven PostgreSQL instance set up and running.

### Installation

1. **Clone the repository**:
   ```sh
   git clone https://github.com/sunilsinghx/restapi-gin.git
   cd restapi-gin
   go mod tidy
   ```
2. **Set up environment variables**:
    DB_URL: <Aiven_Postgres_URL>
    PORT: <Port>

3. **Run the application**:
```sh
   go run main.go

   
    

