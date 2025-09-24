FROM golang:1.25.1-alpine AS base

ARG BUILD_DATE
ARG BUILD_VERSION
ARG PORT

LABEL org.opencontainers.image.authors="Rafael Bruno <rafaelbrunosiqueira@gmail.com>" \
    org.opencontainers.image.title="general_server_go" \
    org.opencontainers.image.description="General server using the language Go and an enterprise level architecture" \
    org.opencontainers.image.created=$BUILD_DATE \
    org.opencontainers.image.version=$BUILD_VERSION

ENV PORT=$PORT

EXPOSE $PORT
RUN addgroup -S appuser && adduser -S appuser -G appuser
USER appuser
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go && chown appuser:appuser /app/main

FROM base AS dev
ENTRYPOINT ["./main"]

FROM dev AS test
RUN go test -v ./...

FROM scratch AS prod
WORKDIR /
COPY --from=dev /app/main /main
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group
USER appuser
ENTRYPOINT ["./main"]
HEALTHCHECK CMD curl -f http://127.0.0.1/ || exit 1