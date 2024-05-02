package app

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"in_mem_storage/internal/app/controller"
	"in_mem_storage/internal/domain/command"
	"in_mem_storage/internal/domain/ratelim"
	"strconv"
	"time"
)

func RateLimRoute(c *fiber.Ctx) error {
	ctx := context.Background()

	cfg, valid := c.Locals("RateLimitConfig").(controller.RateLimitConfig)
	if !valid {
		return internalError(c, errors.New("couldn't find rateLimCfg config"))
	}

	user := c.Query("author")

	limitInSecs, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return badRequest(c, err, "limit must be a number in seconds")
	}
	rLim := ratelim.New(time.Now(), time.Duration(limitInSecs)*time.Second)

	err = controller.RateLimit(ctx, cfg.Log, rLim, user, cfg.RLim)

	return err
}

func CmdRoute(c *fiber.Ctx) error {
	ctx := context.Background()

	cfg, ok := c.Locals("CommandConfig").(controller.CommandConfig)
	if !ok {
		return internalError(c, errors.New("couldn't find CommandConfig config"))
	}

	user := c.Query("author")

	cmdName := c.Query("cmd")
	key := c.Query("key")
	b := c.BodyRaw()
	body := string(b)

	cmd, err := command.NewCommand(cmdName, key, user, body)
	if err != nil {
		return badRequest(c, err, "invalid command")
	}

	res, err := controller.Command(ctx, cfg.Log, user, cmd, cfg.RLim, cfg.Cmd)
	if err != nil {
		return internalError(c, err)
	}

	return c.JSON(res)
}
