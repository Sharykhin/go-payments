FROM golang:1.12 as build-env

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

#ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o api ." -command="./api" --graceful-kill=true -color=true
