# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.19-buster as task-builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v ./cmd/main.go


FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=task-builder /app/main /app/main
COPY --from=task-builder /app/configs /app/configs

CMD ["/app/main"]