# gophers-api

This simple API handle a list of Gophers.

<img src="https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png" alt="yoda-gopher.png" width="300"/> <img src="https://raw.githubusercontent.com/scraly/gophers/main/back-to-the-future-v2.png" alt="back-to-the-future-v2.png" width="300"/>

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
$ go build -o bin/gophers-api internal/main.go

// or 

$ task build
```

### Run app 

``` 
$ go run internal/main.go

// or 

$ task run
```

### Serve Swagger UI 

This will open you browser on Swagger UI

``` 
$ task swagger:serve
```

### Test the API

* Get all Gophers:

```bash
$ curl localhost:8080/gophers
```

Response:

```bash
[{"name":"5th-element","path":"5th-element.png","url":"https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png"},{"name":"yoda-gopher","path":"yoda-gopher.png","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}]
```

* Get a Gopher with the input name

```bash
$ curl localhost:8080/gopher?name=5th-element
```

Response:

```bash
{"name":"5th-element","path":"5th-element.png","url":"https://raw.githubusercontent.com/scraly/gophers/main/5th-element.png"}
```

/!\ Returns a 404 HTTP Error Code if a Gopher have not been found for the given name.

* Add a new Gopher

```
$ curl -X POST localhost:8080/gopher \
   -H "Content-Type: application/json" \
   -d '{"name":"yoda-gopher","path":"yoda-gopher.png","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}'  
```

Response:

```bash
{"name":"yoda-gopher","path":"yoda-gopher.png","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}
```

## Notes

This API use [go-swagger](https://goswagger.io/install.html)
