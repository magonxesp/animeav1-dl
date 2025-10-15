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
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
)

// BaseURL es la URL base del sitio web de animeav1
const BaseURL = "https://animeav1.com"

// ExtractEpisodesLinks obtiene todas las URLs de los episodios de una serie de animeav1.com
func ExtractEpisodesLinks(url string) ([]string, error) {
	var episodeLinks []string
	episodePattern := regexp.MustCompile(`^/media/.+/\d+$`)

	collector := NewCollector()

	collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
		linkPath := element.Attr("href")
		if strings.HasPrefix(linkPath, "/media") && episodePattern.MatchString(linkPath) {
			fullURL := BaseURL + linkPath
			episodeLinks = append(episodeLinks, fullURL)
		}
	})

	err := collector.Visit(url)
	if err != nil {
		return nil, err
	}

	return episodeLinks, nil
}

// ExtractEpisodeDownloadLink obtiene el enlace de descarga de Mega para un episodio específico
func ExtractEpisodeDownloadLink(url string) (string, error) {
	// Crear un contexto de Chrome configurado
	ctx, cancel := NewChromeContext()
	defer cancel()

	var downloadLink string

	// Automatizar la navegación y extracción del enlace
	err := chromedp.Run(ctx,
		// Navegar a la URL
		chromedp.Navigate(url),
		// Esperar a que el botón de descarga esté visible y hacer clic
		chromedp.Click("button[aria-label=\"Descargar\"]", chromedp.NodeVisible),
		// Esperar 100ms para que se procese el JavaScript
		chromedp.Sleep(100*time.Millisecond),
		// Esperar a que aparezca el enlace de Mega y extraer su href
		chromedp.AttributeValue(`//a[.//span[text()="Mega"]]`, "href", &downloadLink, nil, chromedp.NodeVisible),
	)

	if err != nil {
		return "", fmt.Errorf("error al extraer el enlace de descarga: %w", err)
	}

	if downloadLink == "" {
		return "", fmt.Errorf("no se encontró el enlace de descarga de Mega")
	}

	return downloadLink, nil
}

// ExtractEpisodeDownloadLinks obtiene todos los enlaces de descarga de Mega para una serie
func ExtractEpisodeDownloadLinks(url string) ([]string, error) {
	episodeLinks, err := ExtractEpisodesLinks(url)

	if err != nil {
		return nil, fmt.Errorf("error al obtener lista de episodios: %w", err)
	}

	var downloadLinks []string
	fmt.Println("Extrayendo enlaces de descarga:")

	for _, episodeLink := range episodeLinks {		
		megaLink, err := ExtractEpisodeDownloadLink(episodeLink)
		if err != nil {
			fmt.Printf("Error al extraer enlace de descarga para %s: %v\n", episodeLink, err)
			continue
		}

		downloadLinks = append(downloadLinks, megaLink)
	}

	return downloadLinks, nil
}