# Estructura del proyecto

- **cmd**: Aqui iran los comandos de cobra
- **internal**: Aqui ira la logica interna del programa
- **pkg**: Aqui iran paquetes que puedan ser potencialmente genericos e independientes del programa y puedan ser publicados
- **main.go**: Punto de entrada del programa
- **src**: Donde esta el codigo fuente del frontend
- **index.html**: Punto de entrada del frontend

# Tecnologias

Para el cli y web server:
- Go

Para el frontend:
- React
- TypeScrpt
- Vite
- Shadcn UI

# Build

## Cli y Web server

**No usar GOCACHE**

Hay que hacer el build con el siguiente comando:

```sh
go build
```

# Test

Para ejecutar los test:

```sh
go test -v
```

## Frontend

Modo desarollo:

```sh
npm run dev
```

Compilar para produccion:

```sh
npm run build
```

Encontrar errores de codigo estaticos:

```sh
npm run lint
```

Arreglar errores de codigo estaticos y de estilo de codigo:

```sh
npm run lint:fix
```

Para añadir un componente de Shadcn:

```sh
npx shadcn@latest add <nombre del componente en minuscula>
# Por ejemplo
npx shadcn@latest add button
```