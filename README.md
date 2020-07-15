# Go JWT

A quick intro to JWT with Go, based on [this](https://www.udemy.com/course/build-jwt-authenticated-restful-apis-with-golang/).

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

To start Postgres, use the `docker-compose.yml` file. The Go app isn't running in Docker, but Postgres is because I don't have it installed on my computer.

Enter the container like this, if necessary:

```text
$ dc exec -it <container_id> psql -U postgres
```
