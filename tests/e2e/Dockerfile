# Simple usage with a mounted data directory:
# > docker build -t stratos-chain-e2e --build-arg uid=$(id -u) --build-arg gid=$(id -g) -f tests/e2e/Dockerfile .
FROM golang:1.20-alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Install minimum necessary dependencies
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/stratosnet/stratos-chain

# Add source files
COPY . .
RUN make install


# Final image
FROM alpine:edge

ENV WORK_DIR /stchaind
ENV RUN_AS_USER stratos

# Install ca-certificates
RUN apk add --update ca-certificates

ARG uid=2048
ARG gid=2048

RUN addgroup --gid $gid "$RUN_AS_USER" && \
    adduser -S -G "$RUN_AS_USER" --uid $uid "$RUN_AS_USER" -h "$WORK_DIR"

WORKDIR $WORK_DIR

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/stchaind /usr/bin/stchaind

COPY tests/e2e/entrypoint.sh /usr/bin/entrypoint.sh
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT ["/usr/bin/entrypoint.sh"]
CMD ["stchaind"]
