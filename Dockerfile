FROM golang:1.13.3-alpine3.10 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/ilovejs/profile

COPY Gopkg.lock Gopkg.toml ./
COPY vendor vendor
COPY util util
COPY event event
COPY db db
COPY schema schema
COPY profile-service profile-service
#COPY query-service query-service
#COPY pusher-service pusher-service

RUN go install ./...

FROM alpine:3.10
WORKDIR /usr/bin
COPY --from=build /go/bin .
