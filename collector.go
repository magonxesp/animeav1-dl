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
	"context"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
)

// NewCollector crea una nueva instancia de colly.Collector con la configuración por defecto
func NewCollector() *colly.Collector {
	collector := colly.NewCollector()
	collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"
	return collector
}

// NewChromeContext crea un nuevo contexto de chromedp configurado para ejecución headless
func NewChromeContext() (context.Context, context.CancelFunc) {
	// Crear un contexto con timeout de 30 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	// Configurar opciones para modo headless
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-dev-shm-usage", true),
	)

	// Crear un contexto con las opciones de headless
	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	
	// Crear el contexto final de chromedp
	ctx, cancel = chromedp.NewContext(allocCtx)

	return ctx, cancel
}