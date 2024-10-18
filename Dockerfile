FROM golang:1.23.1-alpine3.20 as Builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -a -o tinyurl


FROM golang:1.23.1-alpine3.20 as Runtime

WORKDIR /app

COPY --from=Builder /app/tinyurl ./

CMD ["./tinyurl"]
