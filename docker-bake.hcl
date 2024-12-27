group "default" {
  targets = ["gophers-api"]
}

target "gophers-api" {
  context = "./gophers-api"
  dockerfile = "Dockerfile"
  tags = ["scraly/gophers-api:bake"]
}
