package request

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"in_mem_storage/internal/domain/command"
	"in_mem_storage/internal/domain/locerr"
	"net/http"
)

func InvalidIpError() error {
	return locerr.New("[InvalidIpError] : specified ip address is not valid", 1)
}

func MissingParamError(missing string) error {
	return locerr.New(fmt.Sprintf("[RequestBodyError] : missing required parameter: %v", missing), 1)
}

type BasicRequest http.Request

func (r *BasicRequest) TryInto() (command.Command, error) {
	return command.FromRequest((*http.Request)(r))
}

type FiberRequest struct {
	Ctx *fiber.Ctx
}
