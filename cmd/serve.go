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
package cmd

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/magonxesp/animeav1-dl/internal/common"
	"github.com/spf13/cobra"
)

type (
	linkRequest struct {
		URL string `json:"url"`
	}

	linkResponse struct {
		Links []string `json:"links"`
	}
)

func newServeCmd() *cobra.Command {
	var addr string

	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Inicia un servidor HTTP para exponer la extracción de enlaces",
		RunE: func(cmd *cobra.Command, args []string) error {
			if addr == "" {
				addr = ":8080"
			}

			return runHTTPServer(addr)
		},
	}

	cmd.Flags().StringVar(&addr, "addr", ":8080", "Dirección en la que escuchar (por defecto :8080)")

	return cmd
}

func runHTTPServer(addr string) error {
	router := chi.NewRouter()
	router.Post("/links", makeLinkHandler(common.ExtractEpisodeDownloadLinks))

	server := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	logger.Info("servidor HTTP escuchando", "addr", addr)

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) && err != nil {
		return err
	}

	return nil
}

func makeLinkHandler(extractor func(string) ([]string, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		var request linkRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			writeJSONError(w, http.StatusBadRequest, "cuerpo JSON inválido")
			return
		}

		if err := common.ValidateMediaURL(request.URL); err != nil {
			if errors.Is(err, common.ErrEmptyURL) || errors.Is(err, common.ErrInvalidMediaURL) {
				writeJSONError(w, http.StatusBadRequest, err.Error())
				return
			}

			logger.Error("error inesperado validando la URL", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		links, err := extractor(request.URL)
		if err != nil {
			logger.Error("error al extraer enlaces de descarga", "error", err, "url", request.URL)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		writeJSON(w, http.StatusOK, linkResponse{Links: links})
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Error("error al codificar la respuesta JSON", "error", err)
	}
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
