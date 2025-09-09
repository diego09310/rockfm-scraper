# Rockfm Scraper
Guarda la lista de canciones que suenan en [RockFM](https://www.rockfm.fm/).

## Requisitos previos
- Se requiere `go` para compilar la aplicación.
- Se requiere `ffmpeg` para ejecutar la aplicación.

## Compilar
```bash
go build
```

## Ejecutar
Para guardar en la base de datos
```bash
./rockfmScraper --db
```
Para mostrar en consola
```bash
./rockfmScraper --console
```
Para guardar en fichero
```bash
./rockfmScraper --console >> songs.txt
```

## Docker

Puedes ejecutar la aplicación usando Docker. Solo necesitas el archivo `Dockerfile` y, si prefieres usar Docker Compose, también el archivo `docker-compose.yml`.

### Ejecución con Docker

```bash
docker build -t rockfm-scraper .
docker run --rm -v $(pwd)/data:/app/data rockfm-scraper
```

### Ejecución con Docker Compose

```bash
docker compose up
```

El directorio `data` se montará automáticamente como volumen. Si no existe, Docker lo creará. Puedes modificar la configuración del volumen en `docker-compose.yml` si lo necesitas.
