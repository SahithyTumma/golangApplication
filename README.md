# Golang Application

This project is a Go application of Facts (questioin and an answer) demonstrating CRUD operations using the Fiber framework, PostgreSQL for data storage, and Apache Pulsar for messaging. It uses GORM for ORM and Fiber for HTTP routing.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete `Fact` entities.
- **PostgreSQL Integration**: Persistent data storage using GORM and PostgreSQL.
- **Apache Pulsar**: Integrated for message publishing and consumption.
- **Fiber Framework**: Efficient and fast HTTP framework for Go.


### File Descriptions

- **`cmd/main.go`**: The entry point of the Go application. Sets up the web server, connects to the database, initializes Pulsar, and defines the routes.

- **`database/database.go`**: Contains functions for connecting to and initializing the PostgreSQL database. Sets up the `gorm` ORM and runs migrations.

- **`handlers/facts.go`**: Defines the HTTP handlers for CRUD operations on the `Fact` model. Handles listing, creating, showing, updating, and deleting facts.

- **`models/models.go`**: Defines the data models used in the application.

- **`pulsar/pulsar.go`**: Manages Pulsar client operations. Includes functions for initializing the Pulsar client, publishing and consuming messages, and closing the client.

- **`Dockerfile`**: Contains instructions for building a Docker image of the Go application. Sets up the working directory, copies files, and installs dependencies.

- **`docker-compose.yml`**: Defines Docker Compose configuration for local development. Sets up services for the Go application, PostgreSQL, and Pulsar.

- **`.env`**: Environment variables used by Docker Compose to configure database credentials.

- **`k8s/db-deployment.yaml`**: Kubernetes deployment configuration for the PostgreSQL database.

- **`k8s/db-service.yaml`**: Kubernetes service configuration for exposing the PostgreSQL database.

- **`k8s/web-deployment.yaml`**: Kubernetes deployment configuration for the Go application.

- **`k8s/web-nodeport.yaml`**: Kubernetes service configuration for exposing the Go application via a NodePort.

- **`k8s/configmap.yaml`**: Kubernetes ConfigMap for storing environment variables like database credentials.

## Setup

**Clone the Repository**

   ```bash
   git clone https://github.com/SahithyTumma/golangApplication.git
   cd golangApplication
   ```

### Build and Run with Docker Compose

To build and run the application using Docker Compose:

    ```bash
    docker-compose up --build
    ```

    This command starts the following services:

    - **Web**: The Go application
    - **DB**: PostgreSQL database
    - **Pulsar**: Apache Pulsar

Interact with APIs at `http://localhost:3000`


### Kubernetes Deployment

To deploy the application in a Kubernetes cluster:

Ensure you have a Kubernetes cluster running and Helm installed. Install apache pulsar and make sure all the pulsar pods are either in `Running` or `Completed` state.

1. **Apply Configurations**: Use `kubectl` to apply the deployment and service configurations.

    ```bash
    kubectl create -f configmap.yaml
    kubectl create -f db-deployment.yaml
    kubectl create -f db-service.yaml
    kubectl create -f web-deployment.yaml
    kubectl create -f web-nodeport.yaml
    ```

2. **Access the Application**: Get the NodePort service's external IP and port:

    ```bash
    minikube service web-service-node --url
    ```

Interact with APIs at `http://localhost<NODE_PORT>` of web-service-node.


## API Endpoints

The application exposes the following REST API endpoints:

- **`GET /`**: List all facts
- **`POST /fact`**: Create a new fact
- **`GET /fact/:id`**: Get a fact by ID
- **`PATCH /fact/:id`**: Update a fact by ID
- **`DELETE /fact/:id`**: Delete a fact by ID


