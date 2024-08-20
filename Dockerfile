FROM golang:1.22-alpine

WORKDIR ./app

EXPOSE 8080
CMD ["./bin/api"]

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/api ./cmd/ \
    && go build -o ./bin/migrate ./cmd/migration/