FROM docker.io/golang:1.12.5-alpine3.9

LABEL  maintainer = "Sumukha PK"

WORKDIR github.com/SUMUKHA-PK/GetService

COPY . .

RUN apk add git

RUN go build -v ./cmd/...

EXPOSE 3000

