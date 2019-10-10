FROM golang:1.12 as build-env

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

ARG install_goose

RUN if [ $install_goose = "on" ] ; then go get -u github.com/pressly/goose/cmd/goose ; else echo Argument is $install_goose ; fi

COPY . .

COPY .docker/go/wait.sh /wait.sh

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

#ENTRYPOINT bash /run.sh

#ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o api ." -command="./api" --graceful-kill=true -color=true
