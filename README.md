# Plantilla para GO y HTMX

Plantilla básica para iniciar un proyecto usando Go y HTMX.

### Tecnologías

- Go
- HTMX
- Templ
- Echo
- Air
- Node
- PNPM
- TailwindCSS y DaisyUI

### Requisitos

Este proyecto necesita de que tengas Node, Go y [Templ](https://templ.guide/) instalado en tu sistema.

Además, se recomienda el uso de [Air](https://github.com/air-verse/air) para el desarrollo.

### Instalación

1. Clonar el repositorio:

```sh
git clone https://github.com/Gabo-div/htmx-go-template
```

2. Navegar al directorio del proyecto:

```sh
cd htmx-go-template
```

3. Instalar las dependencias de JavaScript:

```sh
pnpm i
```

4. Instalar las dependencias del Go:

```sh
 go mod tidy
```

### Desarrollo

Dentro del proyecto ejecuta el siguiente comando para levantar el servidor de desarrollo:

```sh
air
```

Probablemente, quieras compilar los estilos de TailwindCSS mientras vas realizando cambios, para realizar esto ejecuta:

```sh
make tailwind-watch
```

Además, también puedes reiniciar el navegador cuando modifiques tus templates usando:

```sh
make templ-watch
```
