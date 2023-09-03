# simple-shop
This project simulates the backend operations of a simple e-commerce platform. It's built using Golang, and leverages gRPC and gRPC-Gateway for efficient, strongly-typed, and easily-scalable API endpoints.

### Technologies Used
Golang: Backend implementation
gRPC: API design and communication
gRPC-Gateway: To expose the gRPC API over HTTP/REST
PostgreSQL: Database
Docker: Containerization and dependency management

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [Postman](https://www.postman.com/downloads/)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install sqlc
    ```

### Database Schema
The database schema is available for public viewing at [Dbdocs](https://dbdocs.io/thaithian1999/simple_shop).

### API Documentation
Once the server is running, the API documentation can be accessed through Swagger UI at http://localhost:8080/swagger.
