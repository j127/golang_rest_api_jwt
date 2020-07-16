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
$ go get -u github.com/davecgh/go-spew/spew
$ go get golang.org/x/crypto/bcrypt
```

During the refactoring section he added a package named `github.com/subosito/gotenv`, but I had already added [github.com/joho/godotenv](https://github.com/joho/godotenv) in an earlier stage, because I didn't want to put secrets directly in the code.

To start Postgres, use the `docker-compose.yml` file. The Go app isn't running in Docker, but Postgres is because I don't have it installed on my computer.

Enter the container like this, if necessary:

```text
$ dc exec -it <container_id> psql -U postgres
```

## Usage

Sign up for an account:

```text
$ curl -XPOST http://localhost:8000/signup -d '{"Email":"alice@example.com","Password":"1234"}'
```

Log in:

```text
$ curl -XPOST http://localhost:8000/login -d '{"Email":"alice@example.com","Password":"1234"}'
```

Take the token that is returned and use it to access the protected route in the next step.

`GET` a protected route:

```text
$ curl -H 'Authorization: bearer <some_token>' http://localhost:8000/protected
```
