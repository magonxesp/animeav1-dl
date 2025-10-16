package common

import (
	"log/slog"
	"os"
)

var Logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
