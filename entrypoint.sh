#!/bin/bash

go mod vendor

go build -o ./bin/antares-api

./bin/antares-api