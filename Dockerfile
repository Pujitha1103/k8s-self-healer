FROM golang:1.20-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o /out/k8s-self-healer ./cmd/controller

FROM alpine:3.18
COPY --from=builder /out/k8s-self-healer /usr/local/bin/k8s-self-healer
ENTRYPOINT ["/usr/local/bin/k8s-self-healer"]
