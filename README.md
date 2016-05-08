# Golang // Postgres // Echo API Starter

This is a starter codebase for building API's in Golang. This includes, Postgres for the database driver and echo for the framework. This example also uses fasthttp. 

Please submit pull requests if you spot anything that could be done better.

A Dockerfile is also included. 

Note: this is mostly for my own convenience, as I write a lot of microservices in Golang. 

## Building

You can build this the easy way `go build`, or `go run main.go`. Or you can use `make`, which builds a linux binary, which you can use with the Docker image.

## Run

Again, you can run it on your host `go run main.go` or `./api-starter` if you've already compiled the binary. Or you can run within Docker `make run` (be sure to have ran `make` first for this). 

## Use
1. Create some dummy data `$ curl -i -XPOST --url http://localhost:7000/api/v1/things -d '{ "title": "This is a test", "amount": 12 }' --header "Content-Type: application/json"` 
2. Check it was created `$ curl -i -XGET --url http://localhost:7000/api/v1/things`
