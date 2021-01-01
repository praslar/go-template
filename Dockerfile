FROM golang:alpine as build-env

ARG VERSION=0.0.0

# cache dependencies first
WORKDIR /go-template
COPY go.mod /go-template
COPY go.sum /go-template
RUN go mod download

# lastly copy source, any change in source will not break above cache
COPY . /go-template

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w -X main.version=${VERSION}" -o /go/bin/go-template ./main.go

# # <- Second step to build minimal image
FROM alpine:3.11

# RUN apk add --no-cache git ca-certificates tzdata

# we have no self-sign certificate, don't need to update
# && update-ca-certificates
WORKDIR /go-template
COPY ./conf/app.conf /go-template/conf/app.conf
COPY ./conf/status.yml /go-template/conf/status.yml
COPY --from=build-env /go/bin/go-template /go/bin/go-template

ENTRYPOINT ["/go/bin/go-template"]
