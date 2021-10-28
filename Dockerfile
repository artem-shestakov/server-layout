# Buid app
FROM golang:1.16-alpine as builder
WORKDIR /go/src/app/

RUN apk add --no-cache git make gcc g++

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build

# Generate swagger docs
FROM quay.io/goswagger/swagger as docs
WORKDIR /docs

COPY . .
RUN unset GOPATH && swagger generate spec -o swagger.yml --scan-models

# Finish image
FROM alpine
WORKDIR /app

COPY --from=builder /go/src/app/api .
COPY --from=docs /docs/swagger.yml ./docs/

EXPOSE 8000
ENTRYPOINT [ "./api" ]
