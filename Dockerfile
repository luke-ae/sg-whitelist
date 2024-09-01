FROM alpine:3.20

ENV STARGAZE_REST_API_URL=http://rest.stargaze-apis.com
ENV CONSTELLATIONS_GRAPHQL_URL=https://constellations-api.mainnet.stargaze-apis.com/graphql
ENV PORT=42069

COPY .env /app/.env
ADD ./build/server /usr/bin/server
RUN apk add -U --no-cache ca-certificates

CMD ["server"]
