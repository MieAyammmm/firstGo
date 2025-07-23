FROM golang:1.23.0

# Set direktori kerja di dalam container
WORKDIR /app


COPY . .

WORKDIR /app/recipe/backend


RUN go mod download


RUN go build -o out .


CMD ["./out"]
