# go-api-docker
A simple demo for learning how to create a containerized Go application using Docker with configure a CI/CD pipeline pipeline.

## build the docker image with Dockerfile
```shell
#build
docker build -t go-api-demo:v0.1 .
#run
docker run -p 8080:8080  go-api-demo:v0.1
```