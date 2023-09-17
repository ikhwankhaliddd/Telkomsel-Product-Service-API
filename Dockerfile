# Tahap Build
FROM golang:1.19 AS build

# Set working directory
WORKDIR /app

# Salin file Go source code Anda ke dalam container
COPY ./cmd/server . 

# Copy go.mod dan go.sum dari direktori proyek Anda
COPY go.mod .
COPY go.sum .

# Build aplikasi Go (tanpa penyalinan go.mod)
RUN go build -o telkomsel-usecase-service-api

# Tahap Runtime
FROM alpine:latest

# Set working directory di dalam container
WORKDIR /app

# Salin binary aplikasi Go dari tahap build ke tahap runtime
COPY --from=build /app/telkomsel-usecase-service-api .

# Expose port yang akan digunakan oleh aplikasi Anda
EXPOSE 9090

# Perintah untuk menjalankan aplikasi ketika container dimulai
CMD ["./telkomsel-usecase-service-api"]
