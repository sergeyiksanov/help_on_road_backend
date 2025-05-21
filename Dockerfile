FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

WORKDIR /app/cmd
RUN go build -o help_on_road

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/cmd/help_on_road .

COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY --from=builder /app/cmd/help_on_road .
COPY --from=builder /app/db/migrations /root/db/migrations

EXPOSE 8000
EXPOSE 8002

CMD goose -dir /root/db/migrations postgres "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB_NAME}?sslmode=disable" up && ./help_on_road
