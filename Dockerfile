ARG GO_VERSION=1.17.6

FROM golang:${GO_VERSION}-alpine AS build

WORKDIR /src
COPY ./ ./

RUN go test -timeout 30s
RUN go build -o /astro-api .

FROM gcr.io/distroless/base AS final

LABEL maintainer="rtovey"
USER nonroot:nonroot

COPY --from=build --chown=nonroot:nonroot /astro-api /astro-api

EXPOSE 8090

ENTRYPOINT ["/astro-api"]