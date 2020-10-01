FROM golang:1.15.2

WORKDIR /go/src/antares-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]