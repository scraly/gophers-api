FROM golang:1.25-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./gophers-api internal/main.go

FROM dhi.io/alpine-base:3.23 AS final
COPY --from=builder /build/gophers-api .
EXPOSE 8080
CMD ["./gophers-api"]
