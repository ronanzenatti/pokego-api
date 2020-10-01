#!/bin/bash

if [ $ENV == "DEV" ]; then
    air -c .air.toml
else
    go build -o ./bin/app

    ./bin/app
fi