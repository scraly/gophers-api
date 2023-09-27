# gophers-api

This simple API handle a list of Gophers.
It alllows to:
- list the existing Gophers
- display the information about a Gopher
- create a new Gopher
- delete a Gopher
- update the path and the URL of a Gopher

<img src="https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png" alt="yoda-gopher.png" width="300"/> <img src="https://raw.githubusercontent.com/scraly/gophers/main/back-to-the-future-v2.png" alt="back-to-the-future-v2.png" width="300"/>

## Gitpod integration

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/scraly/gophers-api.git)

## Docker image

Gophers API is available in [Docker Hub](https://hub.docker.com/r/scraly/gophers-api).

### Run the Gophers API with Docker

```bash
docker run -p 8080:8080 scraly/gophers-api:linux-amd64
```

## How to install 

### Prerequisites

Install Go in 1.16 version minimum.  
Install [Taskfile](https://taskfile.dev/#/installation) (optional):

```bash
brew install go-task/tap/go-task
```

Install go-swagger:

```bash
brew tap go-swagger/go-swagger
brew install go-swagger
swagger version
```

### Build 

``` 
go build -o bin/gophers-api internal/main.go

// or 

task build
```

### Run app 

``` 
go run internal/main.go

// or 

task run
```

### Serve Swagger UI 

This will open you browser on Swagger UI

``` 
task swagger.serve
```

### Test the API

* Get all Gophers:

```bash
curl localhost:8080/gophers
```

Response:

```bash
[{"displayname":"5th Element","name":"5th-element","url":"https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png"}]
```

* Get a Gopher with the input name

```bash
curl "localhost:8080/gopher?name=5th-element"
```

Response:

```bash
{"displayname":"5th Element","name":"5th-element","url":"https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png"}
```

/!\ Returns a 404 HTTP Error Code if a Gopher have not been found for the given name.

* Add a new Gopher

```
curl -X POST localhost:8080/gopher \
   -H "Content-Type: application/json" \
   -d '{"name":"yoda-gopher","displayname":"Yoda Gopher","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}'  
```

Response:

```bash
{"displayname":"Yoda Gopher.png","name":"yoda-gopher","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}
```

Add another Gopher:

```
curl -X POST localhost:8080/gopher \
   -H "Content-Type: application/json" \
   -d '{"name":"jurassic-park","displayname":"Gopher Park","url":"https://raw.githubusercontent.com/scraly/gophers/main/jurassic-park.png"}'  
```

* Delete a Gopher

```bash
curl -X DELETE "localhost:8080/gopher?name=5th-element"
```

* Update a Gopher

```bash
curl -X PUT localhost:8080/gopher \
   -H "Content-Type: application/json" \
   -d '{"name":"yoda-gopher","displayname":"El mejor Yoda Gopher","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}' 
```

Response:

```bash
{"displayname":"El mejor Yoda Gopher","name":"yoda-gopher","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}
```

## Build docker image

* Build a docker image for our current/host platform:

```
DOCKER_BUILDKIT=1 docker build -t gophers-api .
```

* Build for GitPod (linux/amd64) and push to the Docker Hub:

```
docker buildx build --platform linux/amd64 -t scraly/gophers-api:linux-amd64 . --push
```

## GoReleaser

Install the CLI:

```
brew install goreleaser
```

Generate .goreleaser.yml (the firt time):
```
goreleaser init
```

Release:

```
goreleaser release --snapshot --skip-publish --rm-dist
```

## Notes

This API use [go-swagger](https://goswagger.io/install.html)
