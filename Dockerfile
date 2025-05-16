FROM --platform=$BUILDPLATFORM node:18 AS web-builder
WORKDIR /app
COPY . .
RUN make build-web
FROM --platform=$BUILDPLATFORM golang:1.20 AS go-builder
WORKDIR /app
COPY . .
COPY --from=web-builder /app/web/out web/out
RUN make build-go-linux-amd64 build-go-linux-arm64
FROM alpine
ARG USERNAME=app
ARG USER_UID_GID=10000
RUN addgroup -g $USER_UID_GID $USERNAME && adduser -u $USER_UID_GID -G $USERNAME -D $USERNAME
USER $USERNAME
WORKDIR /app
ARG TARGETOS
ARG TARGETARCH
COPY --from=go-builder --chown=$USERNAME:$USERNAME /app/out/$TARGETOS-$TARGETARCH/* .
ENTRYPOINT ["./slimlink"]
EXPOSE 44558
