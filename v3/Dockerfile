##### BUILD STAGE #####
# use BUILDPLATFORM to pin to the native platform to prevent emulation from kicking in
FROM --platform=$BUILDPLATFORM golang:1.23.6-alpine3.21 AS build

# os from --platform linux/amd64
ARG TARGETOS
# architecture from --platform linux/amd64
ARG TARGETARCH
ARG SERVICE_NAME

WORKDIR /app/v3/services/${SERVICE_NAME}
# copy over dependency files and download dependencies
COPY ./v3/services/${SERVICE_NAME}/go.mod ./v3/services/${SERVICE_NAME}/go.sum ./
COPY ./v3/go.mod ./v3/go.sum /app/v3/

RUN go mod download

# copy over source files
COPY . /app

# build the service and output the binary to /tmp/app
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /tmp/app

##### RUNTIME STAGE #####
FROM alpine:3.21.3

# create group and user app
RUN addgroup -S app && adduser -S app -G app

# copy over app binary from build stage
COPY --from=build /tmp/app /usr/local/bin/app

# switch to user app
USER app
WORKDIR /home/app

ENTRYPOINT ["app"]
CMD ["-v=9", "-logtostderr"]
