FROM golang:1.17.5

ARG PROJECT_NAME

WORKDIR /go/src/${PROJECT_NAME}
COPY . .

RUN go mod vendor
RUN go test ./...