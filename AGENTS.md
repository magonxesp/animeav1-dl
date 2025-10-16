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

# Licencia

El programa esta bajo licencia GPLv3 y todos los ficheros de codigo fuente tienen que tener la licencia al principio del fichero:

```text
AnimeAV1-DL - Un programa para extraer enlaces de descarga de animeav1.com
Copyright (C) 2025  MagonxESP

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
```

Para ficheros `.go` se pondra en comentarios con `//`.
En ficheros que cumplan el patron `*.{js,jsx,ts,tsx,css}` se tiene que poner en un doc block.