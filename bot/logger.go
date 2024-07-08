package bot

import (
	"fmt"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func CreateLogger(mode string) (zerolog.Logger, error) {
	var handler io.Writer
	switch mode {
	case "development":
		handler = zerolog.ConsoleWriter{
			Out: os.Stdout,
		}
	case "production":
		handler = &lumberjack.Logger{
			Filename: "./logs/current.log",
			MaxSize:  1,
			MaxAge:   7,
		}
	default:
		invalidMode := fmt.Errorf("failed to create logger: mode to create logger was invalid")
		return *new(zerolog.Logger), invalidMode
	}

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	return zerolog.New(handler).With().Timestamp().Logger(), nil
}
