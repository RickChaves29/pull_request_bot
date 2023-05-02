FROM golang:1.20-alpine3.17

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY internal /app/internal
COPY utils /app/utils
COPY cmd /app/cmd

RUN go build -o bot ./cmd/bot.go

EXPOSE 3030

CMD [ "./bot" ]
