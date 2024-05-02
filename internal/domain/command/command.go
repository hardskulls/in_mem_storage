package command

import (
	"errors"
	"in_mem_storage/internal/domain/record"
	"strings"
)

type Empty struct{}

type Command interface {
	IsCommand()
}

func NewCommand[K comparable, V any](
	cmd string,
	author record.Author,
	k K,
	v V,
) (Command, error) {
	switch strings.ToLower(cmd) {
	case "set":
		return &Set[K, V]{author: author, key: k, value: v}, nil
	case "get":
		return &Get[K]{author: author, key: k}, nil
	case "delete":
		return &Delete[K]{author: author, key: k}, nil
	case "update":
		return &Update[K, V]{author: author, key: k, value: v}, nil
	default:
		return nil, errors.New("not a valid command")
	}
}

type Response[R, C, P any] struct {
	result  R `json:"result"`
	cmd     C `json:"cmd"`
	payload P `json:"payload,omitempty"`
}

func WithPayload[R, C, P any](
	result R,
	cmd C,
	payload P,
) Response[R, C, P] {
	return Response[R, C, P]{
		result:  result,
		cmd:     cmd,
		payload: payload,
	}
}

func NoPayload[R, C any](result R, cmd C) Response[R, C, any] {
	return Response[R, C, any]{
		result:  result,
		cmd:     cmd,
		payload: struct{}{},
	}
}
