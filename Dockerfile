FROM golang:alpine AS build-image

WORKDIR /app

RUN apk --no-cache add git

COPY go.mod go.sum ./
RUN GOPROXY=direct go mod download

COPY /cmd/*.go ./
COPY /pkg ./pkg

RUN CGO_ENABLED=0 go test ./pkg/...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -ldflags="-s -w" main.go


FROM alpine:latest

WORKDIR /app

COPY --from=build-image /app/main ./

ENTRYPOINT ["/app/main"]
