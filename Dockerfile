FROM gcr.io/distroless/base

COPY astro-api /

EXPOSE 8090

ENTRYPOINT ["/astro-api"]