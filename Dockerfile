FROM golang:1.20

LABEL authors="behnam"

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN chmod +x ./entrypoint.sh

RUN go build  -o app cmd/main.go

CMD ["./app"]

ENTRYPOINT ["/app/entrypoint.sh"]
