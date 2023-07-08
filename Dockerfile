FROM golang:1.19-alpine AS build

ARG VERSION

LABEL description="teler WAF reverse proxy tool"
LABEL repository="https://github.com/kitabisa/teler-proxy"
LABEL maintainer="dwisiswant0"

WORKDIR /app
COPY ./go.mod .
RUN go mod download

COPY . .
RUN go build -ldflags "-s -w -X github.com/kitabisa/teler-proxy/common.Version=${VERSION}" \
	-o ./bin/teler-proxy ./cmd/teler-proxy 

FROM alpine:latest

COPY --from=build /app/bin/teler-proxy /bin/teler-proxy
ENV HOME /
ENTRYPOINT ["/bin/teler-proxy"]
