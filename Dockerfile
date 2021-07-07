# ----------------------------------------------------------------------
# Building environment
# ----------------------------------------------------------------------

FROM golang:1.15.8-alpine3.12 AS go-builder

# Set up dependencies
RUN set -eux; apk add --no-cache ca-certificates build-base;
RUN apk add --no-cache make gcc git libc-dev bash linux-headers eudev-dev
# NOTE: add these to run with LEDGER_ENABLED=true
# RUN apk add libusb-dev linux-headers

# Add env variable
ENV GOBIN /go/bin

WORKDIR /code
COPY . /code/


# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v0.14.0-beta3/libwasmvm_muslc.a /lib/libwasmvm_muslc.a
RUN sha256sum /lib/libwasmvm_muslc.a | grep adea8f977601daa8daa9885e02b31ca6dd0ab6d4dbbd8ba2ccfa447ffebda37c

RUN LEDGER_ENABLED=false BUILD_TAGS=muslc make build

# ----------------------------------------------------------------------
# Running environment
# ----------------------------------------------------------------------

FROM ubuntu:18.04

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657
# metrics port
EXPOSE 26660
# grpc port
EXPOSE 9090

#RUN apt update && \
#    apt install -y iputils-ping net-tools vim curl wget musl-dev netcat && \
#    apt clean && apt autoclean

# Bash: konstellation: No such file or directory
# ldd /usr/local/bin/konstellation
# libc.musl-x86_64.so.1 => not found
#RUN ln -s /usr/lib/x86_64-linux-musl/libc.so /lib/libc.musl-x86_64.so.1

COPY --from=go-builder /code/build/knstld /usr/bin/knstld

COPY docker/* /opt/
RUN chmod +x /opt/*.sh

WORKDIR /opt

CMD ["/usr/bin/knstld", "version"]