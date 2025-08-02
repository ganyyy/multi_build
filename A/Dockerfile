# syntax = docker/dockerfile:1.4

ARG GO_VERSION=1.24.5

FROM --platform=${BUILDPLATFORM:-linux/amd64} golang:${GO_VERSION} AS builder


ARG CACHE_GO_MOD_ID="cache-go-mod"
ARG CACHE_GO_BUILD_ID="cache-go-build"

ARG GO_MOD_CACHE="/go-mod-cache"
ARG GO_BUILD_CACHE="/go-build-cache"

ARG BUILD_DIR="/build"


# 必须要输入的构建的服务目录
# 例如：A/SA , A/SB 或者 B/SC, B/SD
# 这个目录下必须有 main.go 文件
# 例如：A/SA/main.go
ARG SERVICE_DIR
ARG NEED_DATA=false

ARG GO_BUILD_FLAGS='-trimpath -ldflags=-s -ldflags=-w -race'

WORKDIR ${BUILD_DIR}

ENV GOCACHE=${GO_BUILD_CACHE}
ENV GOMODCACHE=${GO_MOD_CACHE}
ENV GO111MODULE=on

RUN --mount=type=cache,target=${GO_MOD_CACHE},id=${CACHE_GO_MOD_ID} \
    --mount=type=bind,source=.,target=${BUILD_DIR}/A \
    --mount=type=bind,from=repo_b,target=${BUILD_DIR}/B \
    cd ${SERVICE_DIR} && \
    GOPROXY=https://goproxy.cn,direct go mod download


RUN --mount=type=cache,target=${GO_MOD_CACHE},id=${CACHE_GO_MOD_ID} \
    --mount=type=cache,target=${GO_BUILD_CACHE},id=${CACHE_GO_BUILD_ID} \
    --mount=type=bind,source=.,target=${BUILD_DIR}/A \
    --mount=type=bind,from=repo_b,target=${BUILD_DIR}/B \
    cd ${SERVICE_DIR} && \
    GOOS=linux GOARCH=amd64 go build ${GO_BUILD_FLAGS} -o /out/cmd main.go

RUN --mount=type=bind,source=.,target=${BUILD_DIR}/A \
    --mount=type=bind,from=repo_b,target=${BUILD_DIR}/B \
    mkdir -p /out/data && \
    if [ "${NEED_DATA}" = "true" ]; then cp -a ${SERVICE_DIR}/main.go /out/data/; fi
        

FROM --platform=${BUILDPLATFORM:-linux/amd64} gcr.io/distroless/cc-debian12:nonroot AS runner

WORKDIR /app

COPY --from=builder /out/cmd cmd
COPY --from=builder /out/data data
EXPOSE 8080

CMD ["/app/cmd"]