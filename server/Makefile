BUILD_PREFIX=dist

IMAGE_TAG=latest
IMAGE_NAME=quay.io/hbjy/kube-idp

DOCKER_CMD=docker
DOCKERFILE=deploy/Dockerfile

.PHONY: all

all: clean build

clean:
	rm -rf ${BUILD_PREFIX}/server

build: export CGO_ENABLED = 0
build: export GOOS = linux
build: export GOARCH = amd64
build:
	go build -a -tags netgo -ldflags="-w" -o ${BUILD_PREFIX}/server cmd/main.go

run:
	go run cmd/main.go

docker-build:
	${DOCKER_CMD} build -t ${IMAGE_NAME}:${IMAGE_TAG} -f ${DOCKERFILE} .
