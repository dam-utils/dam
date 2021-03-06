FROM golang:1.14.2

ARG GOOS
ARG GOARCH
ARG PROJECT_NAME
WORKDIR /go/src/${PROJECT_NAME}

COPY . .

RUN go mod vendor
RUN go build -o ${PROJECT_NAME} main.go