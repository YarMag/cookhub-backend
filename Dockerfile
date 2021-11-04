FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .
COPY api ./api
COPY db ./db
COPY assets ./assets

RUN go build -o /cookhub-rest-server

ADD cert.pem /tmp/
ADD key.pem /tmp/
RUN ls /tmp && chmod 777 /tmp/key.pem

EXPOSE 8080
	
ENTRYPOINT ["/cookhub-rest-server"]