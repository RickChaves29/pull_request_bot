FROM golang:1.20-alpine3.17 AS build

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY internal /app/internal
COPY utils /app/utils
COPY cmd /app/cmd

RUN go build -o bot ./cmd/bot.go

FROM alpine:3.17.2 AS prod

WORKDIR /app
COPY --from=build /app/bot /app/
CMD [ "./bot" ]
