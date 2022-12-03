FROM golang:1.19-alpine

WORKDIR /app

RUN apk update && apk add --no-cache alpine-sdk

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .

RUN go build -o /app/token-supply-api

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=0 /app/token-supply-api ./
COPY --from=0 /app/config.yaml ./

EXPOSE 8080

CMD [ "/app/token-supply-api" ]
