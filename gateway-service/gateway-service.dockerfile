FROM golang:1.21.4 as dev


WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download


COPY . .
RUN go build -o /gateway-service
# Expose port 5051

# CMD reflex  -r '\.go$$' -s -- sh -c "go run ./ main.go" 
CMD [ "/gateway-service" ]