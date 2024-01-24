# Pokemon Project

This is a simple project that utilizes PostgreSQL, Go, and React to create a Pokemon application. Please follow the instructions below to set up and run the project locally.

## Setting up the Database

1. Ensure you have PostgreSQL installed on your local machine.

2. Edit the `.env` file in the project root with your local PostgreSQL username and password. Make sure to replace the placeholders with your actual credentials.

    ```env
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=your password
    DB_NAME=pokemon
    ```

3. Create a database named "pokemon" in your local PostgreSQL server.

## Setting up Go (Golang) Backend

1. Install the Go packages by running the following command in the project root:

    ```bash
    go mod download
    ```

## Setting up React Frontend

1. Navigate to the `view/frontend` directory:

    ```bash
    cd view/frontend
    ```

2. Install the required npm packages:

    ```bash
    npm install
    ```

## Running the Application

Enter the `makefile` directory and run the provided scripts.

or

1. To run the backend, execute the following command in the project root:

    ```bash
    go run main.go
    ```

2. To run the frontend, navigate to the `view/frontend` directory and run:

    ```bash
    npm start
    ```
