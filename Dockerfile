# Gunakan gambar base golang
FROM golang:1.19 AS build

# Set working directory
WORKDIR /app

# Salin file Go source code Anda ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o telkomsel-usecase-service-api

# Image container runtime yang lebih ringan
FROM alpine:latest

# Set working directory di dalam container
WORKDIR /app

# Salin binary aplikasi Go dari build stage ke container
COPY --from=build /app/telkomsel-usecase-service-api .

# Expose port yang akan digunakan oleh aplikasi Anda
EXPOSE 9090

# Perintah untuk menjalankan aplikasi ketika container dimulai
CMD ["./telkomsel-usecase-service-api"]
