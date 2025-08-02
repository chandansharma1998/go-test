FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o go-test .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-test .

EXPOSE 8085

CMD ["./go-test"]