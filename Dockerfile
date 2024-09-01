FROM alpine:3.20

COPY .env /app/.env
ADD ./build/server /usr/bin/server
RUN apk add -U --no-cache ca-certificates

CMD ["server"]
