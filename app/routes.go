package app

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"in_mem_storage/internal/app/controller"
	"in_mem_storage/internal/domain/command"
	"in_mem_storage/internal/domain/ratelim"
	"in_mem_storage/internal/domain/ttl"
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

	limitInMS, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return badRequest(c, err, "limit must be a number in milliseconds")
	}
	rLim := ratelim.New(time.Now(), time.Duration(limitInMS)*time.Millisecond)

	err = controller.RateLimit(ctx, rLim, user, cfg)

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
	body := string(c.BodyRaw())

	cmd, err := command.NewCommand(cmdName, key, user, body)
	if err != nil {
		return badRequest(c, err, "invalid command")
	}

	exp, _ := strconv.Atoi(c.Query("expires"))
	expires := time.Now().Add(time.Duration(exp) * time.Second)

	res, err := controller.Command(ctx, cmd, expires, ttl.New(key), cfg)
	if err != nil {
		return internalError(c, err)
	}

	return c.JSON(res)
}
