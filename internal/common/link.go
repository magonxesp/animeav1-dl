// AnimeAV1-DL - Un programa para extraer enlaces de descarga de animeav1.com
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
package common

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
)

const BaseURL = "https://animeav1.com"

func ExtractEpisodesLinks(url string) ([]string, error) {
	var episodeLinks []string
	episodePattern := regexp.MustCompile(`^/media/.+/\d+$`)

	collector := NewCollector()

	collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
		linkPath := element.Attr("href")
		if strings.HasPrefix(linkPath, "/media") && episodePattern.MatchString(linkPath) {
			episodeLinks = append(episodeLinks, BaseURL+linkPath)
		}
	})

	if err := collector.Visit(url); err != nil {
		return nil, err
	}

	return episodeLinks, nil
}

func ExtractEpisodeDownloadLink(url string) (string, error) {
	ctx, cancel := NewChromeContext()
	defer cancel()

	var downloadLink string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Click(`button[aria-label="Descargar"]`, chromedp.NodeVisible),
		chromedp.Sleep(100*time.Millisecond),
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

func ExtractEpisodesDownloadLinks(url string) ([]string, error) {
	episodeLinks, err := ExtractEpisodesLinks(url)
	if err != nil {
		return nil, fmt.Errorf("error al obtener lista de episodios: %w", err)
	}

	var downloadLinks []string

	for _, episodeLink := range episodeLinks {
		megaLink, err := ExtractEpisodeDownloadLink(episodeLink)
		if err != nil {
			return nil, fmt.Errorf("error al extraer enlace de descarga para %s: %w", episodeLink, err)
		}

		downloadLinks = append(downloadLinks, megaLink)
	}

	return downloadLinks, nil
}
