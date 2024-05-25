FROM golang:1.22.3 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/http-mock-service
RUN echo "nobody:x:65534:65534:nobody,,,:/_nonexistent::/bin/false" > /etc/passwd.minimal \
    && echo "nobody:x:65534:nobody" >> /etc/group.minimal

FROM scratch AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/http-mock-service /app/http-mock-service
COPY --from=build-stage /etc/passwd.minimal /etc/passwd
COPY --from=build-stage /etc/group.minimal /etc/group
COPY config.yaml /app/config.yaml

EXPOSE 8080

USER nobody:nobody

WORKDIR /app

ENTRYPOINT ["/app/http-mock-service"]
