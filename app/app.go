package app

import (
	"in_mem_storage/internal/app/server"
)

type App struct {
	server *server.Server
}

func New(port server.Port) *App {
	srv := initServer(port)

	return &App{
		server: srv,
	}
}

func (a *App) Run() error {
	return a.server.Run()
}
