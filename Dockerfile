# build stage
FROM golang:1.12.5-alpine AS build-stage
WORKDIR /src
ADD . /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -extldflags "-static"' -o fridayApp

# final stage
FROM scratch
EXPOSE 80
COPY --from=build-stage /src/templates /templates
COPY --from=build-stage /src/fridayApp /fridayApp
ENTRYPOINT ["/fridayApp"]