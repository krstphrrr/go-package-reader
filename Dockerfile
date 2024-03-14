FROM golang:1.22.1-alpine3.19

WORKDIR /app

COPY ./entrypoint.sh ./
COPY ./go.mod ./
COPY ./src ./


RUN go build -o /bin/app
# CMD ["tail", "/dev/null", "-f"]
# ENTRYPOINT [ "app" ]
ENTRYPOINT [ "/app/entrypoint.sh" ]