FROM golang:1.11-alpine3.7 AS build-env
RUN apk add --no-cache git mercurial musl-dev gcc
WORKDIR /workspace
COPY ./ ./
RUN cd ./cmd && go build -o goapp

# final stage
FROM alpine
WORKDIR /app
RUN apk add --no-cache curl
COPY --from=build-env /workspace/cmd/goapp /app/
ENTRYPOINT ./goapp
