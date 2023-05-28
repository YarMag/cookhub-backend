FROM golang:1.17-alpine

WORKDIR /app

COPY ga_operators ./ga_operators
COPY island_algorithm ./island_algorithm

COPY cache ./cache
COPY db ./db

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .

RUN go build -o /cookhub-ga-optimizer

EXPOSE 8090

ENTRYPOINT ["/cookhub-ga-optimizer"]