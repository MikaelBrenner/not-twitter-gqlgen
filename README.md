
# not-twitter-gqlgen

## Description

A very simple GraphQL-based twitter clone to evaluate the Go [gqlgen](https://github.com/99designs/gqlgen) library. It uses MySQL for data persistence, used without an ORM.

Implemented functionality includes JWT-based authentication, mutations and queries for the `tweet` and `user` entity.

## Comments

I personally did not enjoy the structure created and somewhat enforced by `gqlgen`. While there are many improvements possible to this basic code(externalizing configuration, ORM, Clean Architecture), the schema-first approach, and the generated structure did not appeal to me, which is why I'll probably not be updating this repo.

## Running the API

You can use [docker-compose](https://docs.docker.com/compose/) to run the application:
```bash
docker-compose up -d # the -d option runs the containers in the background
```

This starts the server at the port 9797.


## Usage

The GraphQL Playground(and server) can be accessed at the `/` endpoint, and the available entity is `tweets`, and the available mutations are:

- `createUser`
- `createTweet`
- `login`
- `refreshToken`


# not-twitter-gqlgen
