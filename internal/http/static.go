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
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/magonxesp/animeav1-dl/internal/common"
)

const (
	frontendDirEnv = "ANIMEAV1_FRONTEND_DIR"
	defaultDistDir = "dist"
)

var (
	staticOnce      sync.Once
	staticHandler   http.Handler
	staticDir       string
	staticErr       error
	staticIndexFile string
)

func ensureStaticAssets() {
	staticOnce.Do(func() {
		dir, err := resolveDistDir()
		if err != nil {
			staticErr = err
			return
		}

		if absDir, err := filepath.Abs(dir); err == nil {
			dir = absDir
		}

		staticDir = dir
		staticHandler = http.FileServer(http.Dir(dir))
		staticIndexFile = filepath.Join(dir, "index.html")
	})
}

func resolveDistDir() (string, error) {
	var candidates []string

	if env := os.Getenv(frontendDirEnv); env != "" {
		candidates = append(candidates, env)
	}

	if wd, err := os.Getwd(); err == nil {
		candidates = append(candidates, filepath.Join(wd, defaultDistDir))
	}

	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)
		candidates = append(candidates,
			filepath.Join(exeDir, defaultDistDir),
			filepath.Join(filepath.Dir(exeDir), defaultDistDir),
		)
	}

	seen := make(map[string]struct{})
	for _, candidate := range candidates {
		if candidate == "" {
			continue
		}

		candidate = filepath.Clean(candidate)
		if _, ok := seen[candidate]; ok {
			continue
		}
		seen[candidate] = struct{}{}

		info, err := os.Stat(candidate)
		if err != nil || !info.IsDir() {
			continue
		}

		if _, err := os.Stat(filepath.Join(candidate, "index.html")); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("no se encontró el directorio del frontend; define %s o coloca la carpeta %q junto al binario", frontendDirEnv, defaultDistDir)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	ensureStaticAssets()

	if staticErr != nil {
		common.Logger.Error("frontend no disponible", "error", staticErr)
		http.Error(w, "frontend no disponible", http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, staticIndexFile)
}

func serveStaticAsset(w http.ResponseWriter, r *http.Request) {
	ensureStaticAssets()

	if staticErr != nil {
		http.NotFound(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api/") {
		http.NotFound(w, r)
		return
	}

	cleanedPath := filepath.Clean(strings.TrimPrefix(r.URL.Path, "/"))
	if cleanedPath == "." || cleanedPath == "" {
		serveIndex(w, r)
		return
	}

	fullPath := filepath.Join(staticDir, cleanedPath)

	if rel, err := filepath.Rel(staticDir, fullPath); err != nil || strings.HasPrefix(rel, "..") {
		http.NotFound(w, r)
		return
	}

	info, err := os.Stat(fullPath)
	if err == nil && !info.IsDir() {
		staticHandler.ServeHTTP(w, r)
		return
	}

	serveIndex(w, r)
}
