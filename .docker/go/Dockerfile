FROM golang:1.12 as build-env

ADD . /src

WORKDIR /src

RUN go get -v ./...

RUN cd cmd/api && go build -o api

FROM alpine

WORKDIR /app

COPY --from=build-env /src/cmd/api/api /app/

ENTRYPOINT ./api



