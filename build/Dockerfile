FROM golang:1.17.8 AS builder
WORKDIR /usr/src/requeueip
COPY . .
RUN make

FROM alpine:3.15.0
WORKDIR /
COPY --from=builder /usr/src/requeueip/bin /usr/src/requeueip/bin
COPY ./build/entrypoint.sh /

ENTRYPOINT ["/entrypoint.sh"]