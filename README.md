# Hinge Application

This application exposes endpoints to:
1. Get a user's incoming likes
2. Edit a user's profile

## Deployment

In the root directory, build the docker image: 

`docker build -t hinge-api .`

Bring detached service up with docker-compose:

`docker-compose up -d`

## Authentication

A rudimentary basic auth check is in place that uses the users first name as the username and 'hinge' as the password.
The first name is case sensitive.

`--user Daenerys:hinge`

## Testing

Application is exposed on localhost:8000

### Incoming Likes
For testing incoming likes, try the following commands:

`curl --user Daenerys:hinge "localhost:8000/user/likes?"`

`curl --user Jon:hinge "localhost:8000/user/likes?"`

More users and relationships can be found in the `./db/hinge.sql` file

### Edit Profile

For testing editing a profile, try the following commands:

`curl -vvv -H "Content-Type: application/json" --user Jon:hinge -XPUT "localhost:8000/user/profile" -d '{"last_name": "Stark"}'`

`curl -vvv -H "Content-Type: application/json" --user Jon:hinge -XPUT "localhost:8000/user/profile" -d '{"first_name": ""}'`