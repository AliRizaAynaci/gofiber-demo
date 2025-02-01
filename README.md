# GoFiber Demo Project

This project is a demo project created to learn the Fiber framework in Go. It provides a simple REST API example and interacts with an SQLite database.

## Technology Stack

- **Go**: The programming language used is Go.
- **Fiber**: The web framework used is Fiber.
- **GORM**: The ORM (Object-Relational Mapping) library used is GORM.
- **SQLite**: The database used is SQLite.

## Project Structure

The project includes a basic REST API structure and consists of the following files:

- `main.go`: The main file of the application. The Fiber server is started and routes are defined here.
- `models/`: Directory containing database models.
- `handlers/`: Directory containing API endpoint handlers.
- `database/`: Directory managing database connections and operations.

## Installation

Follow these steps to run the project on your local machine:

1. **Install Go**: If Go is not installed on your computer, download and install it from [Go's official site](https://golang.org/dl/).

2. **Clone the Project**:
   ```bash
   git clone https://github.com/AliRizaAynaci/gofiber-demo.git
   cd gofiber-demo
   ```

3. **Download Dependencies**:
   ```bash
   go mod download
   ```

4. **Prepare the Database**:
   Since SQLite is used in the project, no additional setup is required. The database will be created automatically.

5. **Run the Application**:
   ```bash
   go run main.go
   ```

   The application will run by default at `http://localhost:3000`.

