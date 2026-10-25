FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -ldflags "-X main.version=${VERSION}" -o /app .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
