# build stage
FROM golang:1.24-alpine AS builder

# install ca-certificates for https requests
RUN apk --no-cache add ca-certificates tzdata

# set working directory
WORKDIR /app

# copy go mod files
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# copy source code
COPY . .

# build the application
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o api-center .

# final stage
FROM alpine:latest

# install ca-certificates for https requests
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# copy the binary from builder stage
COPY --from=builder /app/api-center .

# copy locales directory for i18n support
COPY --from=builder /app/locales ./locales
