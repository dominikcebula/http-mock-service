FROM scratch AS build-release-stage

COPY http-mock-service /app/http-mock-service
COPY docker/passwd.minimal /etc/passwd
COPY docker/group.minimal /etc/group
COPY config.yaml /app/config.yaml

EXPOSE 8080

USER nobody:nobody

WORKDIR /app

ENTRYPOINT ["/app/http-mock-service"]
