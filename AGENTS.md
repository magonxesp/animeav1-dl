# Estructura del proyecto

- **cmd**: Aqui iran los comandos de cobra
- **internal**: Aqui ira la logica interna del programa
- **pkg**: Aqui iran paquetes que puedan ser potencialmente genericos e independientes del programa y puedan ser publicados
- **main.go**: Punto de entrada del programa

# Build

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