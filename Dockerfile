FROM golang:1.16-alpine

WORKDIR /app

COPY api ./api
COPY cache ./cache
COPY repositories ./repositories
COPY db ./db
COPY third_party ./third_party
COPY middleware ./middleware
COPY models ./models
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .
COPY assets ./assets

RUN go build -o /cookhub-rest-server

ADD server.crt /tmp/
ADD server.key /tmp/
ADD serviceAccountKey.json /tmp/
RUN chmod 777 /tmp/server.key
RUN chmod 777 /tmp/serviceAccountKey.json

EXPOSE 8080

ENTRYPOINT ["/cookhub-rest-server"]