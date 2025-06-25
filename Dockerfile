FROM golang:1.24.4-alpine3.22 AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/main /main

EXPOSE 8080

CMD ["/main"]
