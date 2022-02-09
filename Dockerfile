FROM golang:1.16 AS build
WORKDIR /app

COPY . .
RUN --mount=type=ssh if [ ! -d "./vendor" ]; then make build.vendor; fi

ARG build_args
RUN GOOS=linux GOARCH=amd64 make build.local


FROM gcr.io/distroless/base-debian11

COPY --from=build /app/target/run /usr/bin/run

# API port
EXPOSE 8080

ENTRYPOINT ["/usr/bin/run"]
