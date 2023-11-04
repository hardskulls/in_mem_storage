package main

import (
	cmdserv "in_mem_storage/internal/application/service/crud_cmd_executor"
	logserv "in_mem_storage/internal/application/service/logger"
	rlimserv "in_mem_storage/internal/application/service/rate_limiter"
	reqserv "in_mem_storage/internal/application/service/server"
	ttlserv "in_mem_storage/internal/application/service/time_to_live"
	req "in_mem_storage/internal/domain/incoming_request/value_object/request"
	resp "in_mem_storage/internal/domain/incoming_request/value_object/response"
	cmdexecrepo "in_mem_storage/internal/infrastructure/db/in_mem/built_in/sync_map/service/command_executor/repository"
	rlimrepo "in_mem_storage/internal/infrastructure/db/in_mem/built_in/sync_map/service/rate_limiter/repository"
	ttlrepo "in_mem_storage/internal/infrastructure/db/in_mem/built_in/sync_map/service/time_to_live/repository"
	logadap "in_mem_storage/internal/infrastructure/io/system/built_in/console/service/logger/adapter"
	reqhandadap "in_mem_storage/internal/infrastructure/net/http/built_in/http_server/service/request_handler/adapter"
	"in_mem_storage/internal/presentation/controllers/net/http/crud_controller"
	"time"
)

const (
	rateLimRoutePath = "/api/v1/rate_limit/"
	crudCmdRoutePath = "/api/v1/db_cmd/"
	interval         = time.Millisecond * 100
)

func main() {
	recRepo := cmdexecrepo.RecordRepo[string]{}
	ttlRepo := ttlrepo.ExpiryRecRepo[time.Time]{}
	rLimRepo := rlimrepo.RateLimitRepo[string]{}

	reqAdapter := reqhandadap.StandardHTTPRequestAdapter{}
	logRecAdapter := logadap.LogRecordAdapter{}

	reqServ := reqserv.New[req.Request, resp.Response](&reqAdapter)
	cmdExServ := cmdserv.New(&recRepo, &ttlRepo)
	rLimServ := rlimserv.New(&rLimRepo)
	ttlServ := ttlserv.New(&ttlRepo)
	logServ := logserv.New(logRecAdapter)

	ctrl := crud_controller.New(reqServ, cmdExServ, rLimServ, ttlServ, logServ)

	rLimRoute := crud_controller.RateLimiterRoute[
		req.Request, string, resp.Response,
	](rateLimRoutePath)
	ttlRoute := crud_controller.TimeToLiveRoute[
		req.Request, resp.Response,
	](interval)
	crudCmdRoute := crud_controller.CrudCommandsRoute[
		string, req.Request, string, resp.Response,
	](crudCmdRoutePath)

	go ctrl.RunConfig(rLimRoute)
	go ctrl.RunConfig(ttlRoute)
	ctrl.RunConfig(crudCmdRoute)
}
