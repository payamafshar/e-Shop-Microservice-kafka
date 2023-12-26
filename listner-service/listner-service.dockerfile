FROM golang:1.21.5-alpine3.18 as dev


WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest
# RUN go build -o /gateway-service
# Expose port 5051
EXPOSE 5000

CMD reflex -g "*.go"  go run ./cmd/api main.go --start-service
# CMD [ "/gateway-service" ]