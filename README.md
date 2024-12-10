# Ordine-API - Order Management

This is an application designed for managing orders (ordines) in restaurants. It is developed in Go, using the Gin and GORM libraries to build an efficient and scalable API.

## Technologies Used

- **Go**: The main language used for the backend.
- **Gin**: A web framework for Go, used to build the API.
- **GORM**: An ORM for Go, used to interact with the database.

## How to Run the Project

This project uses Docker to simplify the development environment and application execution. Follow the steps below to set up and run the application locally.

### 1. Run Migrations

To run the migrations and set up the database, use the following command:

```bash
docker-compose up migrate

```
### 2. Start the Application
To start the API, run:

```bash
docker-compose up api
```

This will start the application server, and you will be able to interact with the API.

### 3. Stop the Application
To stop the application, simply run:

```bash
docker-compose down
```

### 4. Stop the Application and Remove the Database
If you want to stop the application and also remove the database, use:

```bash
docker-compose down
```

Note: Removing the database will require you to run the migrations again.

## Swagger
http://localhost:8080/swagger/index.html

## Docker

how to build the images used in this project and publish them on dockerhub?

### golang
build:
```bash
docker build -t matheushpr9/ordine-api-app:vx.x.x -f build\go\Dockerfile .
```
push:
```bash
docker push matheushpr9/ordine-api-app:v1.0.1
```

### mysql
build:
```bash
docker build -t matheushpr9/ordine-api-database:v1.0.1 -f build\mysql\Dockerfile .
```
push:
```bash
docker push matheushpr9/ordine-api-database:v1.0.1
```

## Contributing
If you would like to contribute to the project, follow these steps:

Fork the repository.
Create a branch for your feature or fix.
Commit your changes.
Open a pull request describing the changes made.
## License
This project is licensed under the  License.