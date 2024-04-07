FROM golang:1.20-alpine as builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./gophers-api internal/main.go

FROM alpine:latest AS final
COPY --from=builder /build/gophers-api .
EXPOSE 8080
CMD ["./gophers-api"]