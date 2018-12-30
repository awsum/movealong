FROM golang:1.11-alpine

WORKDIR /go/src/github.com/awsum/movealong
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /binary .

#
#
#

FROM alpine

ENV CONFIG_FILE=/config.yaml
COPY --from=0 /binary .
RUN apk add --no-cache ca-certificates

ENTRYPOINT /binary --config $CONFIG_FILE
