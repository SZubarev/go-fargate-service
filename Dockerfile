FROM golang:alpine AS build-image

WORKDIR /app

ENV GOPROXY direct
ENV CGO_ENABLED=0

RUN apk --no-cache add git

COPY go.mod go.sum ./
RUN go mod download

COPY /cmd/*.go ./
COPY /pkg ./pkg

RUN go test ./pkg/...

RUN go build -tags lambda.norpc -ldflags="-s -w" main.go


FROM alpine:latest

WORKDIR /app

COPY --from=build-image /app/main ./

ENTRYPOINT ["/app/main"]
