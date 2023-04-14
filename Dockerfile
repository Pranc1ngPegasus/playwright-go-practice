FROM golang:1.20 as builder

WORKDIR /go/src/app

COPY go.* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o app ./cmd/app

FROM ghcr.io/mokmok-dev/playwright-go:main

WORKDIR /

COPY --from=builder /go/src/app/app /app

ENTRYPOINT ["/app"]
