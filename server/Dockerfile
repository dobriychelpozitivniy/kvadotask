FROM golang:1.18.0-alpine3.15 as builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o server ./cmd/main.go


# generate clean, final image for end users
FROM golang:1.18.0-alpine3.15

COPY --from=builder /build/server /app/
COPY --from=builder /build/configs /app/configs/

# executable
ENTRYPOINT [ "/app/server", "--config", "/app/configs/prod" ]
