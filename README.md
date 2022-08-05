# Patika.dev PropertyFinder Go Bootcamp Final Project
##### This repository contains the necessary codes &amp; info about the project given in the Patika.dev PropertyFinder Go bootcamp.

---

# Installation
---

### Prerequisites

- **Docker:** 
  *You can install Docker Desktop by following the instructions on the [Docker Desktop website](https://desktop.docker.com/).*
  
- **PostgreSQL:**
  *You have to run the below script to download and initialize Postgresql*
   ```
    docker run -d --name ctip-pg-cnt -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=kbueRm8nuEj9DBVK -e POSTGRES_DB=postgres -v /Users/tohanilhan/Vault/code/PgDb/DataVolume:/var/lib/postgresql/data -p 5433:5432 postgres:latest
    ```
- **Go:**
  *You have to install Go by following the instructions on the [Go website](https://golang.org/doc/install).*

- **Project:**
  *You have to clone the project from the [Github repository](https://github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project) and run the below script to build the project*
   
   ```bash
    git clone https://github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project.git
    ```

- **Build & Run the Docker Container:**

    *You have to make sure that you are in the same directory as the project*
    ```bash
    docker build -t {your-desired-image-name} .
    ```

     ```bash
   docker run -d --env-file=.env --name {your-desired- container-name} -p {your-desired-port}:8080 {your-desired-image-name}
    ```
   
 - **Or you can run the project directly by running the below script** 
    ```go
    go run main.go
    ```
---