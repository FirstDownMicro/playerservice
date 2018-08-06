FROM golang:alpine as builder
RUN mkdir -p $GOPATH/src/github.com/Nanixel/FirstDownMicro/playerservice
ADD . $GOPATH/src/github.com/Nanixel/FirstDownMicro/playerservice
RUN cd $GOPATH/src/github.com/Nanixel/FirstDownMicro/playerservice && go install

FROM alpine
WORKDIR /app
COPY --from=builder go/bin /app
ENTRYPOINT ./playerservice