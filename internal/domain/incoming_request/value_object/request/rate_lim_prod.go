package request

import (
	errs "in_mem_storage/internal/domain/error/value_object"
	lim "in_mem_storage/internal/domain/rate_limiter/value_object"
	"net"
	"net/url"
	"strconv"
	"time"
)

func InvalidIpError() error {
	return errs.New("[InvalidIpError] : specified ip address is not valid", 1)
}

func (r Request) ProduceRateLim() (lim.RateLimit, error) {
	bodyAsStr, err := r.Body()
	if err != nil {
		return lim.RateLimit{}, errs.FromError(err, 1)
	}

	content, err := url.ParseQuery(bodyAsStr)
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

	return lim.RateLimit{For: forIp, Limit: time.Duration(timeout) * time.Millisecond}, nil
}
