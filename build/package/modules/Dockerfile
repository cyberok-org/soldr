## build binary
FROM golang:1.19 as build-binary

COPY ./build/package/modules /project
WORKDIR /project

RUN ls
RUN go mod vendor \
    && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO111MODULE=on go build \
    -v \
    -mod vendor \
    -o /project/bin/modules \
    /project

## Image
FROM debian:buster as image

RUN mkdir -p /opt/vx_modules/bin

COPY --from=build-binary /project/bin/modules /opt/vx_modules
COPY soldr-modules.json /opt/vx_modules

WORKDIR /opt/vx_modules

RUN apt update
RUN apt install git -y

ENTRYPOINT ["/opt/vx_modules/modules"]
