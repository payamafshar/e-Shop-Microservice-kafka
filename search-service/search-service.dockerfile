
FROM golang:1.21.5 AS dev
ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download


COPY . .
RUN go install github.com/cespare/reflex@latest
# RUN go build -o /gateway-service
# Expose port 5051
EXPOSE 6000

ENTRYPOINT ["make"]
CMD ["dev"]
FROM golang:1.21.5 AS prod
ENV GOPROXY=https://goproxy.io,direct
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin

WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download


COPY . .


EXPOSE ${PORT}


ENTRYPOINT ["make"]
CMD ["run"]



