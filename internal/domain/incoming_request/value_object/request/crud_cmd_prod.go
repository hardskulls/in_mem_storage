package request

import (
	"in_mem_storage/internal/application/service/crud_cmd_executor/abstraction"
	errs "in_mem_storage/internal/domain/error/value_object"
	value_objects2 "in_mem_storage/internal/domain/transaction/command/value_object"
	"net/url"
	"strconv"
	"time"
)

func findKeyVal(content url.Values) (string, string, error) {
	key := content.Get("key")
	if key == "" {
		return "", "", MissingParamError("key")
	}
	value := content.Get("value")
	if value == "" {
		return "", "", MissingParamError("value")
	}
	return key, value, nil
}

func findAndCheckExpiresAfter(content url.Values) (time.Duration, error) {
	expireAfter := content.Get("expires_after")
	if expireAfter == "" {
		return 0, MissingParamError("expire_after")
	}
	expiresAfterInt, err := strconv.ParseUint(expireAfter, 10, 32)
	if err != nil {
		return 0, errs.FromError(err, 1)
	}
	return time.Millisecond * time.Duration(expiresAfterInt), nil
}

func (r Request) ProduceCmd() (abstraction.DefaultCommandExecutor, error) {
	bodyAsStr, err := r.Body()
	if err != nil {
		return nil, errs.FromError(err, 1)
	}

	content, err := url.ParseQuery(bodyAsStr)
	if err != nil {
		return nil, errs.FromError(err, 1)
	}

	command := content.Get("command")
	switch command {
	case "get":
		key := content.Get("key")
		if key == "" {
			return nil, MissingParamError("key")
		}

		return value_objects2.GetCommand{Key: key}, nil
	case "set":
		key, value, err := findKeyVal(content)
		if err != nil {
			return nil, errs.FromError(err, 1)
		}
		expireAfter, err := findAndCheckExpiresAfter(content)
		if err != nil {
			return nil, errs.FromError(err, 1)
		}

		return value_objects2.SetCommand{Key: key, Val: value, ExpiresAfter: expireAfter}, nil
	case "delete":
		key := content.Get("key")
		if key == "" {
			return nil, MissingParamError("key")
		}

		return value_objects2.DeleteCommand{Key: key}, nil
	case "update":
		key, value, err := findKeyVal(content)
		if err != nil {
			return nil, errs.FromError(err, 1)
		}

		return value_objects2.UpdateCommand{Key: key, Val: value}, nil
	default:
		return nil, MissingParamError("command")
	}
}
