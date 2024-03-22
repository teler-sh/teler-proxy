FROM golang:alpine

ARG VERSION="docker"
ARG LDFLAGS="-s -w -X github.com/kitabisa/teler-proxy/common.Version=${VERSION}"
ARG PGO_FILE="default.pgo"

LABEL org.opencontainers.image.authors="Dwi Siswanto <me@dw1.io>"
LABEL org.opencontainers.image.description="teler Proxy enabling seamless integration with teler WAF to protect locally running web service against a variety of web-based attacks"
LABEL org.opencontainers.image.licenses="Apache-2.0"
LABEL org.opencontainers.image.ref.name="${VERSION}"
LABEL org.opencontainers.image.title="teler-proxy"
LABEL org.opencontainers.image.url="https://github.com/kitabisa/teler-proxy"
LABEL org.opencontainers.image.version="${VERSION}"

WORKDIR /app

COPY ["go.mod", "${PGO_FILE}", "./"]
RUN go mod download

COPY . .

ENV CGO_ENABLED=1

RUN apk add build-base
RUN go build \
		-pgo "${PGO_FILE}" \
		-ldflags "${LDFLAGS}" \
		-o /bin/teler-proxy \
		-v ./cmd/teler-proxy

RUN addgroup \
		-g "2000" \
		teler-proxy && \
	adduser \
		-g "teler-proxy" \
		-G "teler-proxy" \
		-u "1000" \
		-h "/app" \
		-D teler-proxy

USER teler-proxy:teler-proxy

ENTRYPOINT ["/bin/teler-proxy"]
