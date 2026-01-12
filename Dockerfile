# -------- BUILD STAGE --------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# install gcc for sqlite
RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o task-service ./cmd/api


# -------- RUNTIME STAGE --------
FROM alpine:latest

RUN apk add --no-cache sqlite

WORKDIR /app

COPY --from=builder /app/task-service .
COPY --from=builder /app/tasks.db ./tasks.db

EXPOSE 8080

CMD ["./task-service"]
