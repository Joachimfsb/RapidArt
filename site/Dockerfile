FROM golang:1.22 as builder

LABEL maintainer="mail@host.tld"
LABEL stage=builder

# Set up execution environment in container's GOPATH
WORKDIR /go/src/app

# Copy relevant folders into container
COPY . /go/src/app/

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o server ./cmd/rapidart/main.go

# Indicate port on which server listens
EXPOSE 8080

# Instantiate binary
CMD ["./server"]
