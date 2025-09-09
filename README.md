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
