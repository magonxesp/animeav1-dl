FROM golang:1.25-alpine AS backend

WORKDIR /build

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o animeav1-dl

FROM node:22-alpine AS frontend

WORKDIR /build

COPY . .

RUN npm install
RUN npm run build

FROM alpine:latest

RUN apk add --no-cache chromium

WORKDIR /app

COPY --from=backend /build/animeav1-dl ./animeav1-dl
COPY --from=frontend /build/dist ./dist

ENV CHROME_BIN=/usr/bin/chromium-browser
ENV CHROME_PATH=/usr/lib/chromium/
ENV ANIMEAV1_FRONTEND_DIR=/app/dist

EXPOSE 8080

ENTRYPOINT ["./animeav1-dl"]
CMD ["serve"]
