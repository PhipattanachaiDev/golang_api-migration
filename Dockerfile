FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o main ./cmd/main.go

EXPOSE 8080
CMD ["./main"]
