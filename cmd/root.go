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
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "animeav1-dl",
		Short: "Extrae enlaces de descarga de animeav1.com",
		Long:  "AnimeAV1 Downloader permite extraer enlaces de descarga de animeav1.com desde la línea de comandos o mediante un endpoint HTTP.",
	}

	logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
)

// Execute ejecuta el comando raíz.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newCLICmd(), newServeCmd())
}
