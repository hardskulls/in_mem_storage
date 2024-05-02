package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

type (
	Port   = int
	Method = string
	Path   = string
)

type Handler func(ctx *fiber.Ctx) error

type Config struct {
	Log *slog.Logger
}

type Server struct {
	port Port
	app  *fiber.App
}

func NewServer(port Port) *Server {
	return &Server{port: port, app: fiber.New()}
}

// Add adds middleware or dependencies.
func (s *Server) Add(m any) {
	_ = s.app.Use(m)
}

func (s *Server) Handle(m Method, p Path, h Handler) {
	s.app.Add(m, p, h)
}

func (s *Server) Run() error {
	return s.app.Listen(fmt.Sprintf(":%d", s.port))
}
