FROM golang:1.19-alpine AS build

ARG VERSION
ARG LDFLAGS="-s -w -X github.com/kitabisa/teler-proxy/common.Version=${VERSION}"

LABEL description="teler Proxy enabling seamless integration with teler WAF to protect locally running web service against a variety of web-based attacks"
LABEL repository="https://github.com/kitabisa/teler-proxy"
LABEL maintainer="dwisiswant0"

WORKDIR /app
COPY ./go.mod .
RUN go mod download

RUN apk add build-base

COPY . .
RUN CGO_ENABLED="1" go build -ldflags "${LDFLAGS}" \
	-o ./bin/teler-proxy ./cmd/teler-proxy 

FROM alpine:latest

COPY --from=build /app/bin/teler-proxy /bin/teler-proxy
ENV HOME /
ENTRYPOINT ["/bin/teler-proxy"]
