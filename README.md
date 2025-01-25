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
docker-compose up app
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
docker-compose down -v
```

Note: Removing the database will require you to run the migrations again.

## Tests
To run the tests use the following command:
```bash
docker-compose up test
```

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

## Swarm  
 iniciando:
```bash
docker swarm init
```
Adcionando outros nós ao cluster
```bash
docker swarm join --token <token> <addres>
```

Adicionar services ao nó baseando-se no docker-compose 
```bash
 docker stack deploy --compose-file docker-compose.swarm.yml ordine
```
vizualizar containers
```bash
docker ps
```
ver logs:
```bash
docker logs ordine_api.1.w2wuo2k000toiyc685hind0yo
```
listar um service especifico
```bash
docker service ps ordine_api
```
escalar horizontalmente
```bash
docker service scale ordine_api=5
```
Fazer o nó atual deixar o cluster
```bash
docker swarm leave --force
```

## Contributing
If you would like to contribute to the project, follow these steps:

Fork the repository.
Create a branch for your feature or fix.
Commit your changes.
Open a pull request describing the changes made.
## License
This project is licensed under the [MIT](https://choosealicense.com/licenses/mit/#) License.