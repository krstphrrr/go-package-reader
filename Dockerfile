FROM golang:1.22.1-alpine3.19

WORKDIR /app

COPY ./entrypoint.sh ./
COPY ./go.mod ./
COPY ./src ./
RUN chmod +x /app/entrypoint.sh

RUN go build -o /bin/app

ENTRYPOINT [ "/bin/sh", "/app/entrypoint.sh" ]