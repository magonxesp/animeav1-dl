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
	"github.com/magonxesp/animeav1-dl/internal/common"
	internalhttp "github.com/magonxesp/animeav1-dl/internal/http"
	"github.com/spf13/cobra"
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

			router := internalhttp.NewRouter()
			server := internalhttp.New(router, internalhttp.WithAddr(addr))
			common.Logger.Info("servidor HTTP escuchando", "addr", addr)

			return server.Listen()
		},
	}

	cmd.Flags().StringVar(&addr, "addr", ":8080", "Dirección en la que escuchar (por defecto :8080)")

	return cmd
}
