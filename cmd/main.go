package main

import (
	cmdserv "in_mem_storage/internal/application/service/crud_cmd_executor"
	logserv "in_mem_storage/internal/application/service/logger"
	rlimserv "in_mem_storage/internal/application/service/rate_limiter"
	reqserv "in_mem_storage/internal/application/service/server"
	ttlserv "in_mem_storage/internal/application/service/time_to_live"
	req "in_mem_storage/internal/domain/incoming_request/value_object/request"
	resp "in_mem_storage/internal/domain/incoming_request/value_object/response"
	log "in_mem_storage/internal/domain/log/value_object"
	logobj "in_mem_storage/internal/domain/log/value_object/log_record"
	stfrup "in_mem_storage/internal/domain/log/value_object/stack_frames_up"
	cmdexecrepo "in_mem_storage/internal/infrastructure/db/in_mem/built_in/sync_map/service/command_executor/repository"
	rlimrepo "in_mem_storage/internal/infrastructure/db/in_mem/built_in/sync_map/service/rate_limiter/repository"
	ttlrepo "in_mem_storage/internal/infrastructure/db/in_mem/built_in/sync_map/service/time_to_live/repository"
	logadap "in_mem_storage/internal/infrastructure/io/system/built_in/console/service/logger/adapter"
	reqhandadap "in_mem_storage/internal/infrastructure/net/http/built_in/http_server/service/request_handler/adapter"
	httpctrl "in_mem_storage/internal/presentation/controllers/net/http/http_controller"
	"os"
	"strconv"
	"time"
)

const (
	rateLimRoutePath           = "/api/v1/rate_limit/"
	crudCmdRoutePath           = "/api/v1/db_cmd/"
	interval                   = time.Millisecond * 100
	serverError      log.Event = "Server error"
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

	ctrl := httpctrl.New(reqServ, cmdExServ, rLimServ, ttlServ, logServ)

	rLimRoute := httpctrl.RateLimiterRoute[
		req.Request, string, resp.Response,
	](rateLimRoutePath)
	ttlRoute := httpctrl.TimeToLiveRoute[
		req.Request, resp.Response,
	](interval)
	crudCmdRoute := httpctrl.CrudCommandsRoute[
		string, req.Request, string, resp.Response,
	](crudCmdRoutePath)

	ctrl.AppendConfigs(rLimRoute, ttlRoute, crudCmdRoute)
	ctrl.RunConfigs(httpctrl.BgGoroutineRunner)

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := ctrl.Run(port)

	logServ.Log(logobj.New(log.Error, serverError, "", err.Error(), stfrup.Here))
}
