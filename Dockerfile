FROM golang:1.25-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

WORKDIR /src

# Install tools needed only at build time
RUN apk add --no-cache ca-certificates

COPY go.mod ./
RUN go mod download && go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -trimpath -ldflags "-s -w" -o /gotter .

RUN chmod +x /gotter

FROM scratch AS release
LABEL org.opencontainers.image.authors="Joseph Rodriguez <josephrodriguez@duck.com>"
LABEL org.opencontainers.image.source="https://github.com/josephrodriguez/gotter"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /gotter /gotter

USER 1000
EXPOSE 8080

ENTRYPOINT ["/gotter"]

