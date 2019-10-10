FROM golang:1.12 as build-env

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

ARG install_goose
ARG install_compile_deamon

RUN if [ $install_goose = "on" ] ; then go get -u github.com/pressly/goose/cmd/goose ; echo "Installing goose is turned off" ; fi
RUN if [ $install_compile_deamon = "on" ] ; then go get github.com/githubnemo/CompileDaemon ; echo "Installing CompileDeamon is turned off" ; fi

COPY . .

COPY .docker/go/rabbitmq.sh /rabbitmq.sh
COPY .docker/go/fluentd.sh /fluentd.sh

#ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o api ." -command="./api" --graceful-kill=true -color=true
