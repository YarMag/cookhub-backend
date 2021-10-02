FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .

RUN go build -o /cookhub-rest-server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /cookhub-rest-server /cookhub-rest-server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/cookhub-rest-server"]