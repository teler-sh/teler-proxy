VERSION  = $(shell git describe --always --tags)

build:
	docker-compose build --build-arg VERSION=${VERSION}

start:
	docker-compose up ${ARGS}

start-daemon: ARGS = -d
start-daemon: start

stop:
	docker-compose down