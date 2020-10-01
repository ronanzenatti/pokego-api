FROM golang:1.15.2

WORKDIR /go/src/project
COPY . .

RUN go get -u github.com/cosmtrek/air
RUN go mod tidy
RUN go mod vendor