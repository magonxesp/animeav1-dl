// AnimeAV1 Downloader - Un programa para extraer enlaces de descarga de animeav1.com
// Copyright (C) 2025  MagonxESP
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Este programa se distribuye con la esperanza de que sea útil,
// pero SIN NINGUNA GARANTÍA; sin siquiera la garantía implícita de
// COMERCIABILIDAD o IDONEIDAD PARA UN PROPÓSITO PARTICULAR.
// Consulte la Licencia Pública General de GNU para más detalles.
//
// Debe haber recibido una copia de la Licencia Pública General de GNU
// junto con este programa.  Si no, consulte <https://www.gnu.org/licenses/>.
package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// NewRouter crea un router de chi con las rutas necesarias.
func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/api/download-links", GetDownloadLinksHandler)
	router.Get("/", serveIndex)
	router.Get("/*", serveStaticAsset)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			http.NotFound(w, r)
			return
		}

		serveIndex(w, r)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	})

	return router
}
