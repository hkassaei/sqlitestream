# Build the order-svc binary
FROM golang:1.19 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source to avoid the need to re-download as well as
# the chance of the downloaded layer to be invalidated due to later changes
RUN go mod download

# Copy the go source
COPY main.go main.go

# Build with CG enabled and with flags needed by go tool to call gcc and build sqlite package statically into the app binary
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -o order-app -ldflags '-linkmode external -extldflags "-static"' main.go

# order-app in the final image
FROM alpine
RUN apk add --no-cache bash
WORKDIR /
COPY --from=builder /workspace/order-app .

ENTRYPOINT ["/order-app"]
