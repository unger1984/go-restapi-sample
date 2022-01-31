FROM golang:1.17 as builder

# Copy go.mod and go.sum separately from the rest of the code,
# so their cached layer is not invalidated when the code changes.
COPY go.mod go.sum /
RUN go mod download

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -trimpath -a -o ./server -ldflags '-extldflags "-static" -s -w' .

# app
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/server .
COPY ./migrations ./migrations
COPY ./config.production.yaml config.yaml

EXPOSE 8080

CMD ["./server", "--config=\"./config.yaml\""]
