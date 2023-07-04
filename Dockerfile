FROM golang:1.20-bookworm as builder

ARG IMG_PATH=/opt/pics
ARG EXHAUST_PATH=/opt/exhaust
RUN apt update && apt install --no-install-recommends libvips-dev -y && mkdir /build
COPY go.mod /build
RUN cd /build && go mod download

COPY . /build
RUN cd /build \
    && go build -ldflags="-s -w" -o gotutu .

FROM debian:bookworm-slim

RUN apt update && apt install --no-install-recommends libvips ca-certificates libjemalloc2 libtcmalloc-minimal4 -y && rm -rf /var/lib/apt/lists/* &&  rm -rf /var/cache/apt/archives/*

COPY --from=builder /build/gotutu    /opt/gotutu
COPY --from=builder /build/data   /opt/data

WORKDIR /opt
VOLUME /opt/data
EXPOSE 3095
CMD ["/opt/gotutu"]