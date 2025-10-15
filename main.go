package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	url := flag.String("url", "", "URL de animeav1.com para extraer los links")
	flag.Parse()

	if *url == "" {
		log.Fatal("El argumento -url es requerido")
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