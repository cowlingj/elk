FROM golang:1.12-alpine3.10 as build

ENV GOPATH=/build/
ENV GOBIN=/build/bin/

COPY /src/main/ /build/src/main
RUN cd /build/src/main
RUN go install main

FROM alpine:3.10.1 as run

EXPOSE 80

COPY --from=build /build/bin/main /

ENV SUCCESS_RATE=0.5
ENTRYPOINT [ "/main" ]