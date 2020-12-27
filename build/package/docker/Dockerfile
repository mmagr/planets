FROM golang:1.15-alpine as builder

WORKDIR /build

# Git must be installed because go mod requires it to download dependencies
# After its used we may remove it to reduce docker image size
RUN apk --no-cache add git tzdata
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o service cmd/weather/api.go

# Execution container
FROM alpine
WORKDIR /app
USER 1000
COPY --from=builder /build/service /app/

CMD ["/app/service"]
