FROM golang:alpine as build
RUN mkdir -p /go/src/github.com/waterborne-labs/instrument-flight-rules
WORKDIR /go/src/github.com/waterborne-labs/instrument-flight-rules
RUN apk update && apk add git
COPY --from=instrumentisto/dep:alpine /usr/local/bin/dep /usr/local/bin/dep
RUN go get -u github.com/gobuffalo/packr/...
COPY . .
RUN dep check
WORKDIR ./cmd
RUN packr
RUN go build -o /ifr .

FROM alpine:latest
WORKDIR /
COPY --from=build /ifr /usr/local/bin/ifr
COPY --from=build /go/src/github.com/waterborne-labs/instrument-flight-rules/ifrs /ifrs

ENTRYPOINT ["ifr"]
