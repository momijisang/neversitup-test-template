FROM golang:1.21-alpine AS build_base

ARG CONFIG_FILE=dev.yml

ENV CGO_ENABLED=1
ENV GO111MODULE=on

RUN apk add --no-cache \
    git gcc g++ curl

# Set the Current Working Directory inside the container
WORKDIR /src

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY ./data/$CONFIG_FILE ./data/config.yml

# Build the Go app
RUN go build -o ./out/app ./cmd/api/main.go


# Start fresh from a smaller image
FROM alpine:3.15
RUN apk add ca-certificates curl tzdata \
    && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build_base /src/out/app /app/restapi
# COPY --from=build_base /src/page /app/page
COPY --from=build_base /src/data /app/data

RUN chmod +x restapi \
    && mkdir log && chmod -R 777 ./log \
    && mkdir tmp && chmod -R 777 ./tmp

# This container exposes port 8080 to the outside world
EXPOSE 3333

HEALTHCHECK --interval=15s --timeout=5s CMD curl -sf http://localhost:3333/api/v1/health-check || exit 1

# Run the binary program produced by `go install`
ENTRYPOINT ./restapi
