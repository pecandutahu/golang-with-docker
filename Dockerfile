# Gunakan image Go sebagai image dasar
FROM golang:1.22.5

# Set environment variables
ENV GO111MODULE=on

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod dan go.sum ke dalam working directory
COPY go.mod go.sum ./

# # Download semua dependensi
RUN go mod download

# Copy seluruh kode aplikasi ke dalam working directory
COPY . .

# Build aplikasi
RUN go build -o main cmd/main.go

# Jalankan aplikasi
CMD ["./main"]
