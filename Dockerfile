FROM golang:1.20-bullseye as builder

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's/security.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt update && apt install --no-install-recommends libvips-dev -y && mkdir /build
COPY go.mod /build
RUN cd /build && go mod download

COPY . /build
RUN cd /build \
    && go build -ldflags="-s -w" -o gotutu .

FROM debian:bullseye-slim

RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's/security.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt update && apt install --no-install-recommends libvips ca-certificates libjemalloc2 libtcmalloc-minimal4 -y && rm -rf /var/lib/apt/lists/* &&  rm -rf /var/cache/apt/archives/*

COPY --from=builder /build/gotutu    /opt/gotutu

WORKDIR /opt
VOLUME /opt/data
EXPOSE 3095
CMD ["/opt/gotutu"]