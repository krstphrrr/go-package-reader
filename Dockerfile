FROM golang:1.22.1-alpine3.19

WORKDIR /app

COPY ./src ./

RUN go build -o /bin/app main.go

ENTRYPOINT [ "app" ]