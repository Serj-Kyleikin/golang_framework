# Stage 1: Builder
FROM golang:1.24.4-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -o server ./Infrastructure/LoadBalancer

# Stage 2: Runtime
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/Infrastructure/LoadBalancer/settings ./settings
ENTRYPOINT ["./server"]