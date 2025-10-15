FROM golang:1.21-alpine

# Instalar Chrome y dependencias necesarias
RUN apk add --no-cache \
    chromium \
    chromium-chromedriver \
    ca-certificates

# Establecer la variable de entorno para chromedp
ENV CHROME_BIN=/usr/bin/chromium-browser
ENV CHROME_PATH=/usr/lib/chromium/

WORKDIR /app

# Copiar los archivos del módulo Go
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o animeav1-dl

# Ejecutar la aplicación
ENTRYPOINT ["./animeav1-dl"]