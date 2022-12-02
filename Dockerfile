FROM golang:1.19-alpine

WORKDIR /app

RUN apk update && apk add alpine-sdk

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN go build -o /token-supply-api

EXPOSE 8080

CMD [ "/token-supply-api" ]
