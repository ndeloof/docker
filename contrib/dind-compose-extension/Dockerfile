# syntax=docker/dockerfile:1

# Demo: docker-in-docker image bundling the whole compose-stack-extension PoC:
#   - dockerd built from this tree, i.e. the extensions framework
#     (moby/moby#53021) plus the in-tree "compose" extension exposing the
#     com.docker.compose.stack.v0 gRPC service on the daemon socket;
#   - a docker compose CLI plugin built from the matching Compose-side PoC
#     (docker/compose#14040), which publishes the resolved stack to that
#     service at the start of "up";
#   - a sample project in /demo.
#
# Build (from the repository root):
#   docker build -f contrib/dind-compose-extension/Dockerfile -t dind-compose-ext .
# Run and exercise:
#   docker run -d --privileged --name dind-compose-ext -e DOCKER_TLS_CERTDIR= dind-compose-ext
#   docker exec -w /demo dind-compose-ext docker compose up
#   docker logs dind-compose-ext 2>&1 | grep "compose extension"
#   -> level=info msg="compose extension: received compose stack definition" project=demo complete=true ...

FROM golang:1.26-alpine AS build
RUN apk add --no-cache git
WORKDIR /src
COPY . .
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -tags "exclude_graphdriver_btrfs" -o /out/dockerd ./cmd/dockerd

FROM golang:1.26-alpine AS compose-build
RUN apk add --no-cache git
ARG COMPOSE_REPO=https://github.com/docker/compose.git
ARG COMPOSE_REF=poc-compose-stack-extension
RUN git clone --depth 1 -b ${COMPOSE_REF} ${COMPOSE_REPO} /src
WORKDIR /src
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o /out/docker-compose ./cmd

FROM docker:28-dind
COPY --from=build /out/dockerd /usr/local/bin/dockerd
COPY --from=compose-build /out/docker-compose /usr/local/libexec/docker/cli-plugins/docker-compose
COPY contrib/dind-compose-extension/demo/compose.yaml /demo/compose.yaml
