FROM golang:1.14.2

ARG PROJECT_NAME

WORKDIR /go/src/${PROJECT_NAME}
COPY . .

RUN go mod vendor
RUN go test -v ./...