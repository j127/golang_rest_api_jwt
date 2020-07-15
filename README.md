# Golang: Intro to REST API JWT auth with Go

From a [Udemy course](https://www.udemy.com/course/build-jwt-authenticated-restful-apis-with-golang/).

## The Application

 | Endpoint     | Handler Function    | HTTP Action |
 |--------------|---------------------|-------------|
 | `/signup`    | `signup`            | `POST`      |
 | `/login`     | `login`             | `POST`      |
 | `/protected` | `protectedEndpoint` | `GET`       |

## Dependencies

```text
$ go get -u github.com/gorilla/mux
$ go get -u github.com/dgrijalva/jwt-go
$ go get -u github.com/lib/pq
```

To start Postgres:

```
$ docker run --rm \
    -p 5432:5432 \
    -v $HOME/docker/volumes/postgres:/var/run/postgresql \
    -e POSTGRES_USER=pv \
    -e POSTGRES_PASSWORD=docker \
    --name postgres-go postgres
```
