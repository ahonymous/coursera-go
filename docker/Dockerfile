FROM golang:latest as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

COPY ./go.mod /app/
RUN go mod download
COPY . /app

RUN go build -o /usr/local/bin/coursera main.go

CMD ["air"]
