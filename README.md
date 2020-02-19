# Hinge Application

This application exposes endpoints to:
1. Get a user's incoming likes
2. Edit a user's profile

## Deployment

Build the docker image 

`docker build -t hinge-api .`

Bring detached service up with docker-compose

`docker-compose up -d`

Port exposed on localhost:8000

## Testing

