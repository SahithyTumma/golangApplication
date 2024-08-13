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

Ensure you have a Kubernetes cluster running and Helm installed. Install apache pulsar and make sure all the pulsar pods are either in `Running` or `Completed` state.

1. **Deploy PostgreSQL**

    ```bash
    kubectl create -f db-deployment.yaml
    kubectl create -f db-service.yaml

2. **Deploy Golang application**

   ```bash
   kubectl create -f web-deployment.yaml
   kubectl create -f web-nodeport.yaml

Access the application at http://<NODE_IP>:<NODE_PORT> of web-service-node NodePort

Interact with API using curl, Postman or Thunderclient in vscode.

## Endpoints

- **GET** `/` - List all facts
- **POST** `/fact` - Create a new fact
- **GET** `/fact/:id` - Get a fact by ID
- **PATCH** `/fact/:id` - Update a fact by ID
- **DELETE** `/fact/:id` - Delete a fact by ID

