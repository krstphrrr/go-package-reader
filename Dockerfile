FROM golang:1.22.1-alpine3.19

WORKDIR /app

COPY ./go.mod ./
COPY ./src ./

# CMD ["tail", "/dev/null", "-f"]
RUN go build -o /bin/app

ENTRYPOINT [ "app" ]