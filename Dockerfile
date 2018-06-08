FROM golang:latest as builder
WORKDIR /go
COPY hack.go .
RUN go build hack.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/hack .
RUN chmod +x hack
ENTRYPOINT ["./hack"]
