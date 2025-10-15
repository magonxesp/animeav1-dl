FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY . .

RUN go mod download
RUN go build -o animeav1-dl

FROM alpine:latest

RUN apk add --no-cache chromium

COPY --from=builder /app/animeav1-dl /usr/local/bin/animeav1-dl

ENV CHROME_BIN=/usr/bin/chromium-browser
ENV CHROME_PATH=/usr/lib/chromium/

ENTRYPOINT ["/usr/local/bin/animeav1-dl"]