# container image that runs your code
FROM gcr.io/distroless/base

# copies your code file from your action repository to the filesystem path `/` of the container
COPY --chown=nonroot:nonroot ./build/plant-service /app/plant-service

USER nonroot:nonroot

WORKDIR /app

ENTRYPOINT ["/app/plant-service"]