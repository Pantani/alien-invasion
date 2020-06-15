FROM golang:alpine as builder
RUN mkdir /build
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o bin/invasion ./cmd

FROM alpine:latest
COPY --from=builder /build/bin /bin/
COPY --from=builder /build/test /bin/test/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/bin/invasion", "-w", "/bin/test/world_4.txt", "-i", "10000", "-a", "10"]
