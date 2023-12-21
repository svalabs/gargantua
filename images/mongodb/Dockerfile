##### sdk image #####
FROM golang:1.21 AS sdk

WORKDIR /go/src/github.com/hobbyfarm/gargantua/services/mongodbsvc
COPY . .

ENV GOOS=linux
ENV CGO_ENABLED=0

RUN go get -d -v ./...
RUN go install -v ./...

#RUN ls -lart && go build -o /go/bin/mongodb main.go
###### release image #####
FROM alpine:latest

COPY --from=sdk /go/bin/mongodbsvc /usr/local/bin/

ENTRYPOINT ["mongodbsvc"]
CMD ["-v=9", "-logtostderr"]