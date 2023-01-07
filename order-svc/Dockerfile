# Build the order-svc binary
FROM golang:1.19 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go

# Build
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -o order-app -ldflags '-linkmode external -extldflags "-static"' main.go

# order-app in the final image
FROM alpine
RUN apk add --no-cache bash
WORKDIR /
COPY --from=builder /workspace/order-app .

ENTRYPOINT ["/order-app"]
