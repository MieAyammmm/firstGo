FROM golang:1.23.0

# Buat workdir di dalam container
WORKDIR /app

# Salin hanya folder backend (karena go.mod di dalamnya)
COPY backend/ .

# Jalankan go mod download
RUN go mod download

# Build
RUN go build -o main .

# Jalankan
CMD ["./main"]
