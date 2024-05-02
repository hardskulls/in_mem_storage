package command

import (
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"strings"
)

func FromRequest(r *http.Request) (Command, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	path := r.URL.Path
	query := r.URL.Query()

	cmdName, _, _ := strings.Cut(path, "/")
	author := query.Get("author")
	key := query.Get("key")
	value := string(body)

	cmd, err := NewCommand(cmdName, author, key, value)

	return cmd, err
}

func FromFiberCtx(c *fiber.Ctx) (Command, error) {
	body := c.BodyRaw()

	cmdName := c.Query("cmd")
	author := c.Query("author")
	key := c.Query("key")
	value := string(body)

	cmd, err := NewCommand(cmdName, author, key, value)

	return cmd, err
}
