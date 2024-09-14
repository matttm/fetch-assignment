# Fetch Assignment

This repository is an API for receipt processot.

This is a demonstration of ability in Go.

## Prerequesites

Docker

## Running

First build the image. When you're in root of the repo, run
```
docker build -t fetch .
```
The image's name is fetch. Next, you can start a container from the built image by running
```
docker run -P fetch
```
The `-P` flag publishes all exposed ports. To find which host port is mapped to the container's exposed port, run
```
docker ps
```
for a result like
```
CONTAINER ID   IMAGE     COMMAND    CREATED          STATUS          PORTS                     NAMES
949e5dd02e8a   fetch     "./main"   33 seconds ago   Up 32 seconds   0.0.0.0:55000->8080/tc
```
This indicates that port 55000 on localhost is mapped to 8080 of the container, so be sure to use the host port when trying to interact with the containrer from the host. For instance, I used the address `localhost:55000` when working in Postman.
