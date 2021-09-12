package app

import (
	"os"
	//"os/signal"
)

type (
	App struct {
		fd *os.File
	}
)

func NewApp() *App {
	return &App{}
}
