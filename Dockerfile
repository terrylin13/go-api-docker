FROM golang:1.18.2

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./bin/api ./cmd/api/api.go

EXPOSE 8080

ENTRYPOINT [ "./bin/api" ]