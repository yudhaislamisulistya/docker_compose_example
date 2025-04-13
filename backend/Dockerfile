# Gunakan image Golang sebagai base
FROM golang:1.18-alpine

# Set working directory di dalam container
WORKDIR /app

# Salin go.mod dan go.sum untuk menginstall dependensi
COPY go.mod go.sum ./
RUN go mod download

# Salin semua file aplikasi
COPY . .

# Build aplikasi
RUN go build -o main .

EXPOSE 1323

# Tentukan perintah untuk menjalankan aplikasi
CMD ["./main"]
