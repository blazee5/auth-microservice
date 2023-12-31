FROM golang:1.21

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

EXPOSE 4000

CMD ["./cmd/main"]