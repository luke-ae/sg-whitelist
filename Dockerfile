FROM alpine:3.20

ADD ./build/server /usr/bin/server
RUN apk add -U --no-cache ca-certificates

CMD ["server"]
