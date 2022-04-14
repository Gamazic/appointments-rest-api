FROM golang:1.18 as builder

WORKDIR /app

COPY ./go.sum ./go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/appointments-rest-api ./appointments-rest-api

EXPOSE 8080
ENTRYPOINT ["./appointments-rest-api"]
