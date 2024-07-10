package bot

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/rs/zerolog"
)

type App struct {
	Client bot.Client
	Config Config
	Logger zerolog.Logger
}

func New(logger zerolog.Logger, config Config) *App {
	return &App{
		Config: config,
		Logger: logger,
	}
}

func (app *App) SetupApp () {}