#!/bin/bash

if [ $APP_ENV == "DEV" ]; then
    go get -u github.com/cosmtrek/air
fi

go mod tidy
go mod vendor

if [ $APP_ENV == "DEV" ]; then
    air -c .air.toml
else
    go build -o ./bin/app
    ./bin/app
fi