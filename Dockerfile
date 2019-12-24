FROM golang:1.13-alpine as builder

RUN apk update && apk add git

WORKDIR /go/src/gitlab.com/EDSN/prototype/diva-go-backend

COPY . .

ENV GO111MODULE off

RUN go get -u github.com/golang/dep/cmd/dep && dep ensure && go build -o /tmp/diva-go-backend

FROM alpine:3.10

COPY --from=builder /tmp/diva-go-backend /

ENTRYPOINT /diva-go-backend
