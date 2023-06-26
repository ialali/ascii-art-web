#!/bin/bash

IMAGE_NAME="ascii-art-web-image"

# Build the Docker image
docker image build -f Dockerfile -t ascii-art-web-image .

#Run the container
 docker container run -p 5050 --detach --name ascii-web-container ascii-web-image