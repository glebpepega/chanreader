package logger

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

func New() (log *slog.Logger) {
	log = slog.New(tint.NewHandler(
		os.Stdout,
		&tint.Options{
			AddSource:  true,
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}))

	return
}
