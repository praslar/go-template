FROM golang:alpine as build-env

ARG VERSION=0.0.0

# cache dependencies first
WORKDIR /gitlab.com/jamalex-ltd/r-d-department/go-template
COPY go.mod /gitlab.com/jamalex-ltd/r-d-department/go-template
COPY go.sum /gitlab.com/jamalex-ltd/r-d-department/go-template
RUN go mod download

# lastly copy source, any change in source will not break above cache
COPY . /gitlab.com/jamalex-ltd/r-d-department/go-template

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w -X main.version=${VERSION}" -o /go/bin/gitlab.com/jamalex-ltd/r-d-department/go-template ./main.go

# # <- Second step to build minimal image
FROM alpine:3.11

# RUN apk add --no-cache git ca-certificates tzdata

# we have no self-sign certificate, don't need to update
# && update-ca-certificates
WORKDIR /gitlab.com/jamalex-ltd/r-d-department/go-template
COPY ./conf/app.conf /gitlab.com/jamalex-ltd/r-d-department/go-template/conf/app.conf
COPY ./conf/status.yml /gitlab.com/jamalex-ltd/r-d-department/go-template/conf/status.yml
COPY --from=build-env /go/bin/gitlab.com/jamalex-ltd/r-d-department/go-template /go/bin/gitlab.com/jamalex-ltd/r-d-department/go-template

ENTRYPOINT ["/go/bin/gitlab.com/jamalex-ltd/r-d-department/go-template"]
