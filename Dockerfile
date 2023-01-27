FROM golang:alpine

WORKDIR /SMS

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /SMS ./...

EXPOSE 40080

ENV API_PORT="40080"

ENTRYPOINT ["/SMS/social-media-api"]