# Golang Application

This project is a Go application demonstrating CRUD operations using the Fiber framework, PostgreSQL for data storage, and Apache Pulsar for messaging. It uses GORM for ORM and Fiber for HTTP routing.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete `Fact` entities.
- **PostgreSQL Integration**: Persistent data storage using GORM and PostgreSQL.
- **Apache Pulsar**: Integrated for message publishing and consumption.
- **Fiber Framework**: Efficient and fast HTTP framework for Go.


## Setup

### Local Development with Docker

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/golangApplication.git
   cd golangApplication

2. **Build and Run with Docker Compose**

    ```bash
    docker compose up --build

    This command starts the following services:

- **Web**: The Go application
- **DB**: PostgreSQL database
- **Pulsar**: Apache Pulsar (standalone mode)



### Kubernetes Deployment

1. **Apply the yaml files for PostgreSQL**

    ```bash
    kubectl apply -f db-deployment.yaml
    kubectl apply -f db-service.yaml
