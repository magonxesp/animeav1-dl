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
package cmd

import (
	"fmt"
	"strings"

	"github.com/magonxesp/animeav1-dl/internal/common"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "animeav1-dl",
	Short: "Extrae enlaces de descarga de animeav1.com",
	Long:  "AnimeAV1 Downloader permite extraer enlaces de descarga de animeav1.com desde la línea de comandos o mediante un endpoint HTTP.",
	RunE:  runRoot,
}

var (
	mediaURL     string
	logDirectory string
	logJSON      bool
	logLevel	 string
)

// Execute ejecuta el comando raíz.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	
	cobra.OnInitialize(initLogging)
	setupLoggingFlags()
	rootCmd.Flags().StringVar(&mediaURL, "url", "", "URL de animeav1.com para extraer los links")
	rootCmd.AddCommand(newServeCmd())
}

func setupLoggingFlags() {
	defaultLogDir := common.GetLogDirectory()
	defaultLogJSON := common.GetLogJSON()
	defaultLogLevel := strings.ToLower(common.GetLogLevel().String())

	logDirectory = defaultLogDir
	logJSON = defaultLogJSON
	logLevel = defaultLogLevel

	rootCmd.PersistentFlags().StringVar(&logDirectory, "log-directory", defaultLogDir, "Directorio donde guardar el fichero de logs")
	rootCmd.PersistentFlags().BoolVar(&logJSON, "log-json", defaultLogJSON, "Emitir los logs en formato JSON")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", defaultLogLevel, "Nivel de los mensajes de logs, puede ser: info, warn, error o debug")
}

func initLogging() {
	common.SetLogDirectory(logDirectory)
	common.SetLogJSON(logJSON)
	common.SetLogLevel(logLevel)
	common.ConfigureLogger()
}

func runRoot(cmd *cobra.Command, args []string) error {
	if err := common.ValidateMediaURL(mediaURL); err != nil {
		return err
	}

	downloadLinks, err := common.ExtractEpisodesDownloadLinks(mediaURL)
	if err != nil {
		return fmt.Errorf("error al extraer los enlaces de descarga: %w", err)
	}

	fmt.Println("\nEnlaces de descarga encontrados:")
	for _, link := range downloadLinks {
		fmt.Println(link)
	}

	return nil
}
