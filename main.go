// AnimeAV1 Downloader - Un programa para extraer enlaces de descarga de animeav1.com
// Copyright (C) 2025  MagonxESP
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

func main() {
	url := flag.String("url", "", "URL de animeav1.com para extraer los links")
	flag.Parse()

	if *url == "" {
		log.Fatal("El argumento -url es requerido")
	}

	matched, err := regexp.MatchString(`^https://animeav1\.com/media/[^/]+`, *url)
	if err != nil {
		log.Fatal("Error al validar la URL:", err)
	}
	if !matched {
		fmt.Println("La URL proporcionada no es válida. Debe seguir el formato: https://animeav1.com/media/<nombre>")
		return
	}

	downloadLinks, err := ExtractEpisodeDownloadLinks(*url)
	if err != nil {
		log.Fatal("Error al extraer los enlaces de descarga:", err)
	}

	fmt.Println("\nEnlaces de descarga encontrados:")
	for _, link := range downloadLinks {
		fmt.Println(link)
	}
}