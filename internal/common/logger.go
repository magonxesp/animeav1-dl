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
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	logDirEnv  	= "LOG_DIRECTORY"
	logJSONEnv 	= "LOG_JSON"
	logLevelEnv = "LOG_LEVEL"
)

var (
	Logger       *slog.Logger
	logDirectory string 		= ""
	logJSON      bool 			= false
	logLevel     slog.Level 	= slog.LevelInfo
)

func ConfigureLogger() {
	Logger = newLogger()
}

func newLogger() *slog.Logger {
	logDir := GetLogDirectory()
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
	options := &slog.HandlerOptions{
		AddSource: true,
		Level:     GetLogLevel(),
	}

	logJson := GetLogJSON()
	if logJson {
		return slog.NewJSONHandler(writer, options)
	}

	return slog.NewTextHandler(writer, options)
}

func init() {
	ConfigureLogger()
}

// SetLogDirectory sobrescribe el directorio de logs (se prioriza LOG_DIRECTORY).
func SetLogDirectory(dir string) {
	logDirectory = dir
}

// SetLogJSON define si el logging debe emitirse en formato JSON (se prioriza LOG_JSON).
func SetLogJSON(enabled bool) {
	logJSON = enabled
}

func SetLogLevel(level string) {
	logLevel = parseLogLevel(level)
}

// GetLogDirectory devuelve el directorio configurado, priorizando la variable de entorno.
func GetLogDirectory() string {
	if value := strings.TrimSpace(os.Getenv(logDirEnv)); value != "" {
		return value
	}

	return logDirectory
}

// GetLogJSON indica si el logging debe emitirse en formato JSON, priorizando la variable de entorno.
func GetLogJSON() bool {
	if value := os.Getenv(logJSONEnv); value != "" {
		return parseBool(value)
	}

	return logJSON
}

func GetLogLevel() slog.Level {
	if value := os.Getenv(logLevelEnv); value != "" {
		return parseLogLevel(value)
	}

	return logLevel
}

func parseLogLevel(level string) slog.Level {
	 switch strings.ToLower(strings.TrimSpace(level)) {
		case "debug":
			return slog.LevelDebug
		case "warn":
			return slog.LevelWarn
		case "error":
			return slog.LevelError
		default:
			return slog.LevelInfo
	}
}