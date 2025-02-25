# Build the application from source
FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV DATABASE_URL="postgresql://admin:admin@localhost:5432/cashflow"

RUN go build -o cashflow ./main.go

EXPOSE 1323

ENTRYPOINT ["/app/cashflow"]