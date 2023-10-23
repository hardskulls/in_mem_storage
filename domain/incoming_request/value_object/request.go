package value_objects

import (
	"fmt"
	errs "in_mem_storage/domain/error"
	lim "in_mem_storage/domain/rate_limiter/value_object"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func MissingParamError(missing string) errs.Error {
	return errs.New(fmt.Sprintf("[RequestBodyError] : missing required parameter: %v", missing), 1)
}

func InvalidIpError() errs.Error {
	return errs.New(fmt.Sprintf("[InvalidIpError] : specified ip address is not valid"), 1)
}

type Request http.Request

func (r *Request) Produce() (lim.RateLimit, errs.Error) {
	var buff []byte
	_, err := r.Body.Read(buff)
	if err != nil {
		return lim.RateLimit{}, errs.FromError(err, 1)
	}

	content, err := url.ParseQuery(string(buff))
	if err != nil {
		return lim.RateLimit{}, errs.FromError(err, 1)
	}

	forIp, limit := content.Get("for_ip"), content.Get("limit")
	if forIp == "" {
		return lim.RateLimit{}, MissingParamError("for_ip")
	}
	if limit == "" {
		return lim.RateLimit{}, MissingParamError("limit")
	}

	ip := net.ParseIP(forIp)
	if ip == nil {
		return lim.RateLimit{}, InvalidIpError()
	}
	timeout, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		return lim.RateLimit{}, errs.FromError(err, 1)
	}

	return lim.RateLimit{For: forIp, Limit: time.Duration(timeout) * time.Millisecond}, errs.Error{}
}
