# Shorter

## Criteria

- URL shortener has 2 APIs, please follow API example to implement:
  - A RESTful API to upload a URL with its expired date and response with a shorten URL.
  - An API to serve shorten URLs responded by upload API, and redirect to original URL. If URL is expired,
    please response with status 404.

## Quick Start

### Golang

- Workdir is short folder.
- Build Golang
  - ```go build -o short .```
- Run Server
  - ```chmod 777 short```
  - ```./short```

Or

- Run Golang
  - ```go run main.go```

### Docker

- Workdir is root folder.
- Build docker:
  - ```docker build . -t short```
- Run docker:
  - ```docker run -p 80:80 short```

---

## Requirement

- Golang version
    - Golang 1.17+
- 3rd party lib 
  - gin-gonic/gin
    - Is popular framework with golang.
    - Have many feature and middleware.
  - chenyahui/gin-cache
    - A high performance gin middleware to cache http response.
    - https://www.cyhone.com/articles/gin-cache/
  - gin-contrib/cors
    - CORS middleware.
  - gin-contrib/pprof 
    - pprof with gin middleware.
  - mattn/go-sqlite3
    - sqlite3 driver.
  - google/uuid
    - generate uuid for short url.
  - spf13/viper 
    - parse environment variables.
- Database
  - SQLite.
    - Lightweight SQL.
    - small, fast, self-contained, high-reliability, full-featured, SQL database engine.
- Unit test lib
  - DATA-DOG/go-sqlmock
  - stretchr/testify/assert

## Feature

- Support
  - Http 2.0.
  - CORS.
  - Pprof and Trace.
- Has cache middleware with memory.
- Use SQLite, not install any sql database.
  - If need to change any RDBMS, Can very quick change.
- CI/CD with gitlab.
- Dockerize.

## TOOD

- Add Singleflight.
  - Prevent Cache Avalanche．
  - https://www.readfog.com/a/1641708239556022272
