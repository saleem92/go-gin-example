FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY ./gogin-example ./gogin-example

COPY ./config ./config

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/gogin-example"]
