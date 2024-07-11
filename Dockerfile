# Gunakan base image Golang
FROM golang:1.18-alpine

# Set environment variables
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

# Install dependencies
RUN apk add --no-cache bash curl

# Install golang-migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz | tar xvz && mv migrate.linux-amd64 /usr/local/bin/migrate

# Buat direktori kerja di dalam container
WORKDIR /app

# Copy semua file ke dalam direktori kerja
COPY . .

# Download dependencies
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Jalankan migrasi dan aplikasi
CMD migrate -database "postgres://postgres:12345678@db:5432/bioskuy_test?sslmode=disable" -path app/migrations up && ./main
