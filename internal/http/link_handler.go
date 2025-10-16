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
package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/magonxesp/animeav1-dl/internal/common"
)

type LinkRequest struct {
	URL string `json:"url"`
}

type LinkResponse struct {
	Links []string `json:"links"`
}

func GetDownloadLinksHandler(writter http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var request LinkRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		RespondJSONError(writter, http.StatusBadRequest, "cuerpo JSON inválido")
		return
	}

	if err := common.ValidateMediaURL(request.URL); err != nil {
		if errors.Is(err, common.ErrEmptyURL) || errors.Is(err, common.ErrInvalidMediaURL) {
			RespondJSONError(writter, http.StatusBadRequest, err.Error())
			return
		}

		common.Logger.Error("error inesperado validando la URL", "error", err)
		writter.WriteHeader(http.StatusInternalServerError)
		return
	}

	links, err := common.ExtractEpisodesDownloadLinks(request.URL)
	if err != nil {
		common.Logger.Error("error al extraer enlaces de descarga", "error", err, "url", request.URL)
		writter.WriteHeader(http.StatusInternalServerError)
		return
	}

	RespondJSON(writter, http.StatusOK, LinkResponse{Links: links})
}
