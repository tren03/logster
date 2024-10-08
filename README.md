# Logster

## Introduction
A toy project done to explore concurrency in go by simulating heavy load on a webserver, whose main function is to upload json sent in post http requests to a blob storage (emulated using azurite docker image)

## Details
The test I ran hit the webserver with 100,000 requests in 7.919 seconds, approx - 12,628.41 requests per second. 
The server was able to handle the load and succesfully upload all data into the blob storage.
We verify this by sending a get request to the webserver after the upload, which logs the size of all blobs and returns total size of all blobs in the container ( each json obj when uploaded is 59 bytes, and when we check the total size, its 5900000 bytes -> this means all 100000 requests were processed wihtout any loss in data

## Demo
### Size of a single request

![Logsize-comment](https://github.com/user-attachments/assets/15c02901-e0de-4b86-8931-38f2c81e247c)

### Handling 100,000 requests (approx - 12,628.41 requests per second)
https://github.com/user-attachments/assets/e124cc5e-dd35-41c2-b40a-ab1f41ffc0d5

1 request = 59 Bytes
100000 request = 5900000 Bytes

## High Level Design
### Description
This Docker Compose setup defines a multi-service architecture with two primary components:
#### 1. Azurite (Blob Storage Emulator):
Azurite is a Docker container that emulates Azure Blob Storage. It allows the application to store data as if it's interacting with real Azure storage.
The container is exposed to other services via an internal Docker network (logster-network) for communication.

#### 2. Logster (Go Application):
Logster is the primary application, which is a Go-based server designed to handle incoming HTTP requests (listening on port 8080). It processes JSON payloads and stores them in Azurite Blob Storage.
The application is built from source using a Dockerfile and waits for Azurite to be ready before starting (depends_on ensures the order).
Environment variables, like AZURE_BLOB_ENDPOINT, allow Logster to communicate with Azurite over the internal Docker network using the service name azurite.

#### Network and Communication:
Both services are connected via a custom bridge network (logster-network), which enables Logster to communicate with Azurite directly using its container name (azurite) instead of an IP address.

![high_level](https://github.com/user-attachments/assets/1b6de655-5f27-4264-99ab-845e09dfe1d9)

## Low Level Design
![low_level](https://github.com/user-attachments/assets/30076bff-1d3f-430b-8b10-d8acc2904594)

## Usage

The application is fully dockerized, making it simple to set up and run.

### Steps to Start

1. Clone the repository:

    ```bash
    git clone https://github.com/tren03/logster.git
    cd logster
    ```

2. Build the Docker images:

    ```bash
    sudo docker-compose build
    ```

3. Run the containers:

    ```bash
    sudo docker-compose up
    ```

### Simulating Load

Make sure you have `ab` (Apache Benchmark) installed:

  ```bash
  sudo apt-get install apache2-utils
  ```

To simulate sending 100,000 requests to the server:
  
  ```bash
  ab -p data.json -T 'application/json' -c 500 -n 100000 http://localhost:8080/log
  ```

To log the size of blobs in azure container after the benchmark

  ```bash
  curl localhost:8080
  ```


