FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-here

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /go-here /go-here

EXPOSE 8087

CMD ["/go-here"]
