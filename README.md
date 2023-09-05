# Simple Shop
This is my portfolio project that simulates the backend operations of a simple e-commerce platform. It's built using Golang, and leverages gRPC and gRPC-Gateway for efficient, strongly-typed, and easily-scalable API endpoints.

### Technologies Used
Golang: Backend implementation <br />
gRPC: API design and communication <br />
gRPC-Gateway: To expose the gRPC API over HTTP/REST <br />
PostgreSQL: Database <br />
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

### Installing
Clone the repository
```bash
git clone https://github.com/theme-ce/simple-shop.git
```

Navigate to the project directory
```bash
cd simple-shop
```

### Setting up the Database
Start the PostgreSQL Docker container
```bash
make postgres
```

Create the database
```bash
make createdb
```

Run migrations
```bash
make migrateup
```

Running the Server
Generate SQL and Protobuf Code
```bash
make sqlc proto
```

Start the server
```bash
make server
```

Test
```bash
make test
```

### Database Schema
The database schema is available for public viewing at [Dbdocs](https://dbdocs.io/thaithian1999/simple_shop).

### API Documentation
Once the server is running, the API documentation can be accessed through Swagger UI at http://localhost:8080/swagger.
