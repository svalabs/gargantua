VERSION 0.8

# target for fetching go dependencies
deps:
    FROM golang:1.21.13-alpine3.20
    WORKDIR /app
    COPY go.mod go.sum ./
    RUN go mod download
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# target to copy over source files
# currently only main.go
artifact:
    FROM +deps
    COPY main.go ./

# target for building the go-binary
# supports goos (default linux) and goarch (default amd64)
build:
    FROM +artifact
    ARG GOOS='linux'
    ARG GOARCH='amd64'
    ARG CGO_ENABLED=0
    RUN CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o /output/app
    SAVE ARTIFACT /output/app

# target to update the dependencies in go.mod
update-deps:
    FROM +artifact
    RUN go get -u ./...
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

# acts as docker base image
# it has a dedicated user and group called app (default user)
# the go-binary is executed from /home/app/app
docker-base:
    ARG TARGETPLATFORM # like linux/arm64, linux/amd64...
    FROM --platform=$TARGETPLATFORM alpine:3.21.2
    RUN addgroup -S app && adduser -S app -G app
    USER app
    WORKDIR /home/app
    ENTRYPOINT ["/home/app/app"]
    CMD ["-v=9", "-logtostderr"]

# builds the docker image containing the go-binary which is placed at /home/app/app
docker:
    ARG TARGETARCH # like arm64, amd64... (earthly built-in)
    ARG TARGETPLATFORM # like linux/arm64, linux/amd64... (earthly built-in)
    FROM --platform=$TARGETPLATFORM +docker-base
    COPY (+build/app --GOARCH=$TARGETARCH) /home/app/app
    SAVE IMAGE gargantua:latest

# builds the docker image for linux/amd64 and linux/arm64
docker-all-platform:
    # parallel build execution
    BUILD --platform=linux/amd64 --platform=linux/arm64/v8 +docker