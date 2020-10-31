# ANTARES-API

Is an example API made entirely in Golang using the Echo framework.

# Installation and running

There are 2 ways available to install and run this project: with Golang installed and without Golang installed.

## With Golang

With Golang, there are two ways to run the API: simply by compiling with Golang or via the Air package.

### **Simply compiling natively**

**Requisites:** 
* [Go 1.15+](https://golang.org/)

**Steps:**
1. Clone this repo.
2. CD at the root of the project.
3. Run:  `go mod vendor`.
4. Create a `.env` file using `.env.example` as example.
5. Run:  `go run main.go`.

### **Using Air package**

**Requisites:**
* [Go 1.15+](https://golang.org/)
* [Air package](https://github.com/cosmtrek/air)

**Steps:**
1. Clone this repo.
2. CD at the root of the project.
3. Run:  `go mod vendor`.
4. Create a `.env` file using `.env.example` as example.
5. Run: `air -c .air.toml`.

--------------

## Without Golang

Without Golang, there are one way to run the API: through the docker.

**Requisites:** 
* [Docker](https://www.docker.com/)

**Steps:**
1. Clone this repo.
2. CD at the root of the project.
3. Run: `docker-compose up`.