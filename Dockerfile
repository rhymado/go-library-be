# Dockerfile digunakan untuk membuat images
# Base Images
FROM golang:1.21.4-alpine

# Environment tambahan (biasanya khusus docker)
ENV GO_ENV=DOCKER DB_HOST=host.docker.internal

# Menentukan working directory di container
WORKDIR /app

# Copy project ke working directory
COPY . .

# Jalankan perintah (instalasi, build, dll) di container
# 1. Install Dependency
RUN go mod download
# 2. Build Application
RUN go build -v -o /app/lib-be ./cmd/main.go

# Expose port
EXPOSE 8080

# Daftarkan aplikasi
ENTRYPOINT [ "/app/lib-be" ]