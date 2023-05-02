FROM golang:1.20-alpine as builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV PORT 8000
ENV PROJECT_DIR /app

WORKDIR $PROJECT_DIR

COPY . .

RUN go build -ldflags="-w -s" ./cmd/orders

FROM scratch

# Copy our static executable.
COPY --from=builder $PROJECT_DIR/app $PROJECT_DIR/app
# Run the hello binary.
ENTRYPOINT ["$PROJECT_DIR/app"]