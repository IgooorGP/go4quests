FROM golang:alpine AS start-point

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . /build

# Copy and download dependency using go mod
RUN go mod download

# Compiles the application to /dist folder with a main binary
RUN go build -o /server /build/main.go

# Build a small image
FROM start-point

COPY --from=start-point /server /
