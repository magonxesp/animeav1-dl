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
package http

import (
	"errors"
	"time"

	nethttp "net/http"

	"github.com/go-chi/chi/v5"
)

// Server encapsula la configuración del servidor HTTP.
type Server struct {
	addr            string
	router          *chi.Mux
	readHeaderLimit time.Duration
}

// Option define una función de configuración para Server.
type Option func(*Server)

// WithAddr establece la dirección de escucha.
func WithAddr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

// WithReadHeaderTimeout permite ajustar el timeout de lectura de cabeceras.
func WithReadHeaderTimeout(d time.Duration) Option {
	return func(s *Server) {
		s.readHeaderLimit = d
	}
}

// New crea un servidor configurado con las opciones especificadas y con el router proporcionado.
func New(router *chi.Mux, opts ...Option) *Server {
	server := &Server{
		addr:            ":8080",
		router:          router,
		readHeaderLimit: 5 * time.Second,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

// Listen inicia el servidor HTTP y devuelve cualquier error inesperado.
func (s *Server) Listen() error {
	httpServer := &nethttp.Server{
		Addr:              s.addr,
		Handler:           s.router,
		ReadHeaderTimeout: s.readHeaderLimit,
	}

	if err := httpServer.ListenAndServe(); !errors.Is(err, nethttp.ErrServerClosed) && err != nil {
		return err
	}

	return nil
}
