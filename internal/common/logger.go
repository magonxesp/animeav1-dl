// AnimeAV1 Downloader - Un programa para extraer enlaces de descarga de animeav1.com
// Copyright (C) 2025  MagonxESP
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Este programa se distribuye con la esperanza de que sea útil,
// pero SIN NINGUNA GARANTÍA; sin siquiera la garantía implícita de
// COMERCIABILIDAD o IDONEIDAD PARA UN PROPÓSITO PARTICULAR.
// Consulte la Licencia Pública General de GNU para más detalles.
//
// Debe haber recibido una copia de la Licencia Pública General de GNU
// junto con este programa.  Si no, consulte <https://www.gnu.org/licenses/>.
package common

import (
	"log/slog"
	"os"
)

var Logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
