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
	"fmt"

	"github.com/magonxesp/animeav1-dl/internal/common"
	"github.com/spf13/cobra"
)

func newCLICmd() *cobra.Command {
	var mediaURL string

	cmd := &cobra.Command{
		Use:   "cli",
		Short: "Extrae enlaces desde la línea de comandos",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := common.ValidateMediaURL(mediaURL); err != nil {
				return err
			}

			downloadLinks, err := common.ExtractEpisodeDownloadLinks(mediaURL)
			if err != nil {
				return fmt.Errorf("error al extraer los enlaces de descarga: %w", err)
			}

			fmt.Println("\nEnlaces de descarga encontrados:")
			for _, link := range downloadLinks {
				fmt.Println(link)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&mediaURL, "url", "", "URL de animeav1.com para extraer los links")
	cobra.CheckErr(cmd.MarkFlagRequired("url"))

	return cmd
}
