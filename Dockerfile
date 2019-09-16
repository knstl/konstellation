# Simple usage with a mounted data directory:
# > docker build -t konstellation .
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.konstellation:/root/.konstellation -v ~/.konstellationcli:/root/.konstellationcli konstellation konstellation init
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.konstellation:/root/.konstellation -v ~/.konstellationcli:/root/.konstellationcli konstellation konstellation start

# ----------------------------------------------------------------------
# Building environment
# ----------------------------------------------------------------------

FROM golang:alpine as build

# Set up dependencies
ENV PACKAGES make gcc git libc-dev bash linux-headers eudev-dev

# Add env variable
ENV GOBIN /go/bin

# Create source directory
RUN mkdir -p /go/src/github.com/konstellation/konstellation

# Add source files
COPY . /go/src/github.com/konstellation/konstellation

# Set working directory for the build
WORKDIR /go/src/github.com/konstellation/konstellation

# Install minimum necessary dependencies, run unit tests
RUN apk add --no-cache $PACKAGES
RUN make tools
RUN make install
#RUN make test_unit

# ----------------------------------------------------------------------
# Running environment
# ----------------------------------------------------------------------

FROM ubuntu:18.04

# p2p port
EXPOSE 26656
# rpc port
EXPOSE 26657
# metrics port
EXPOSE 26660

RUN apt update && \
    apt install -y iputils-ping net-tools vim curl wget musl-dev netcat && \
    apt clean && apt autoclean

# Bash: konstellation: No such file or directory
# ldd /usr/local/bin/konstellation
# libc.musl-x86_64.so.1 => not found
RUN ln -s /usr/lib/x86_64-linux-musl/libc.so /lib/libc.musl-x86_64.so.1

# Copy over binaries from the build
COPY --from=build /go/bin/konstellation         /usr/local/bin/
COPY --from=build /go/bin/konstellationcli      /usr/local/bin/
COPY --from=build /go/src/github.com/konstellation/konstellation/docker/config.toml      /root/.konstellation/
COPY --from=build /go/src/github.com/konstellation/konstellation/docker/genesis.json      /root/.konstellation/

# Init environment
ADD docker/start.sh     /
RUN chmod +x            /start.sh
ADD docker/shutcut/*    /usr/local/bin/
RUN chmod +x            /usr/local/bin/*

WORKDIR /root

# Run daemon
CMD /start.sh