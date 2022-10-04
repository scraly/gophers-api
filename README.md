# gophers-api


* Get all gophers

```bash
GET
https://8080-scraly-gophersapi-bnc6bz9apib.ws-eu67.gitpod.io/gophers

$ curl localhost:8080/gophers
```

* Get all Gophers with the input name

```bash
GET
https://8080-scraly-gophersapi-bnc6bz9apib.ws-eu67.gitpod.io/gophers?name=5th-element

$ curl localhost:8080/gophers?name=5th-element
```

* Add a new Gopher

```
POST 
https://8080-scraly-gophersapi-bnc6bz9apib.ws-eu67.gitpod.io/gopher

$ curl -X POST localhost:8080/gopher \
   -H "Content-Type: application/json" \
   -d '{"name":"yoda-gopher","path":"yoda-gopher.png","url":"https://raw.githubusercontent.com/scraly/gophers/main/yoda-gopher.png"}'  
```