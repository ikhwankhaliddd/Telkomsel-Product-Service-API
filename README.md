
# Product Service API (Telkomsel Usecase Test)

Project ini adalah implementasi dari sebuah backend API service yang menyediakan informasi detail produk. API ini memungkinkan pengguna untuk melakukan operasi Create, Read, Update, dan Delete terhadap produk. Informasi produk yang tersedia mencakup ID produk, nama produk, deskripsi produk, harga produk, variasi produk, rating produk, stok produk, dan gambar produk.

## Tech Stack

**Programming Language:** Golang

**Framework:** Echo

**Database:** PostgreSQL

**Other:** Docker, AWS S3


## Installation

Install Taskfile for better experience.

For MacOS or Linux, you can use this command

```bash
brew install go-task/tap/go-task
```

For Windows, you can use this command

```bash
choco install go-task
```
## Run Locally

1. Clone the project

```bash
  git clone https://github.com/ikhwankhaliddd/Telkomsel-Product-Service-API.git
```

2. Go to the project directory

```bash
  cd Telkomsel-Product-Service-API
```

3. Change the .env.example into .env, and fill the value

```bash
DB_HOST=<YOUR_DB_HOST>
DB_PORT=<YOUR_DB_PORT>
DB_USER=<YOUR_DB_USER>
DB_PASSWORD=<YOUR_DB_PASSWORD>
DB_NAME=<YOUR_DB_NAME>
```

4. Run the migrations

```bash
  go run cmd/migrate/main.go
```

5. Start the server

```bash
  go run main.go
```

Or

You can easily run it with Docker

1. Pull the Docker image from Docker Hub

```bash
docker pull ikhwankhalid/telkomsel-usecase-service-api
```

2. Run the docker image

```bash
docker run -d -p 8080:8080 ikhwankhalid/telkomsel-usecase-service-api
```
## Running Tests

To run tests, run the following command

```bash
  task test:unit
```

Or

```bash
  go test -short  ./...
```


## Documentation

[API Documentation](https://documenter.getpostman.com/view/11005206/2s9YC7Sre9)

