FROM golang:alpine3.21 AS builder

WORKDIR /app

ADD go.mod .

COPY ./ ./

RUN go build -o major ./cmd/main.go

FROM alpine:3.22.0

RUN apk update && \
    apk add --no-cache curl

RUN rm -rf /var/cache/apk/*

WORKDIR /major-app

COPY --from=builder /app/configs/.config.json /major-app/configs/.config.json
COPY --from=builder /app/major /major-app/major
COPY --from=builder /app/docs /major-app/docs

CMD [ "./major" ]