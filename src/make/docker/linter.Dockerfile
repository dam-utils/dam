FROM golangci/golangci-lint:v1.43.0

ARG PROJECT_NAME

WORKDIR /go/src/${PROJECT_NAME}
COPY . .

RUN go mod vendor
RUN golangci-lint run