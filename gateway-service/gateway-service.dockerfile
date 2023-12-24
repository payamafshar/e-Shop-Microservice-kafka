FROM golang:1.21.5 as dev



WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest

# Expose port 5050
EXPOSE 5051

CMD reflex  -r '\.go$$' -s -- sh -c "go run ./cmd/api main.go" 