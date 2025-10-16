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
	"io"
	"log"
	"log/slog"
	"os"
	"path"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	logDirEnv  = "LOG_DIRECTORY"
	logJSONEnv = "LOG_JSON"
)

var Logger = newLogger()

func newLogger() *slog.Logger {
	logDir := os.Getenv(logDirEnv)
	if logDir != "" {
		handler, err := fileHandler(logDir)
		if err != nil {
			panic(err)
		}

		return slog.New(handler)
	}

	return slog.New(consoleHandler())
}

func consoleHandler() slog.Handler {
	return newHandlerFromWritter(os.Stdout)
}

func fileHandler(logDir string) (slog.Handler, error) {
	if _, err := os.Stat(logDir); err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			log.Panicf("can't create %s directory: %v", logDir, err)
		}
	}

	logRotator := &lumberjack.Logger{
		Filename:   path.Join(logDir, "animeav1-dl.log"),
		MaxSize:    100,   // Max size in MB
		MaxBackups: 5,     // Number of backups
		MaxAge:     7,     // Days
		Compress:   true,  // Enable compression
	}

	writter := io.MultiWriter(os.Stdout, logRotator)
	return newHandlerFromWritter(writter), nil
}

func newHandlerFromWritter(writer io.Writer) slog.Handler {
	jsonFormat := parseBoolEnv(logJSONEnv)
	options := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}

	if jsonFormat {
		return slog.NewJSONHandler(writer, options)
	}

	return slog.NewTextHandler(writer, options)
}


