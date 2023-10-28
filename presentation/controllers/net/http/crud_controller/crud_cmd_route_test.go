package crud_controller

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	cmdserv "in_mem_storage/application/service/crud_cmd_executor"
	cmdexec "in_mem_storage/application/service/crud_cmd_executor/abstraction"
	rlimserv "in_mem_storage/application/service/rate_limiter"
	reqserv "in_mem_storage/application/service/server"
	ttlserv "in_mem_storage/application/service/time_to_live"
	reqhdnl "in_mem_storage/domain/incoming_request/value_object"
	rlim "in_mem_storage/domain/rate_limiter/value_object"
	cmds "in_mem_storage/domain/transaction/command/value_object"
	"in_mem_storage/infrastructure/db/in_mem/built_in/sunc_map/service/command_executor/repository"
	repository2 "in_mem_storage/infrastructure/db/in_mem/built_in/sunc_map/service/rate_limiter/repository"
	repository3 "in_mem_storage/infrastructure/db/in_mem/built_in/sunc_map/service/time_to_live/repository"

	//crudctrl "in_mem_storage/presentation/controllers/net/http/crud_controller"
	"sync"
	"testing"
	"time"
)

var storage = ReqRespStorageManualMock{
	Reader: CrudCmdProducerManualMock{},
	Writer: WriterManualMock{data: make([]string, 0)},
}
var user = "user_12345"
var rateLimit = rlim.RateLimit{
	For:      user,
	LastUsed: time.Now(),
	Limit:    time.Millisecond * 500,
}
var result = fmt.Sprintf("[RateLimitOperationSuccess] Your rate limit is, %v", rateLimit)
var command = cmds.SetCommand{
	Key:          "key",
	Val:          "value",
	ExpiresAfter: time.Second * 50,
}

// Request.
type CrudCmdProducerManualMock struct{}

func (r CrudCmdProducerManualMock) ProduceCmd() (cmdexec.DefaultCommandExecutor, error) {
	return command, nil
}

func (r CrudCmdProducerManualMock) Body() string {
	return "body"
}

func (r CrudCmdProducerManualMock) From() string {
	return user
}

func (r CrudCmdProducerManualMock) Date() time.Time {
	return time.Now()
}

// Response.
type WriterManualMock struct {
	sync.Mutex
	data []string
}

func (m *WriterManualMock) Write(str string) error {
	m.Lock()
	m.data = append(m.data, str)
	m.Unlock()
	return nil
}

// Request & response storage.
type ReqRespStorageManualMock struct {
	Reader CrudCmdProducerManualMock
	Writer WriterManualMock
}

// Request handler.
type ReqHandlerManualMock = reqhdnl.ReqHandler[CrudCmdProducerManualMock, *WriterManualMock]

type ReqHandlerPortManualMock struct{}

func (r ReqHandlerPortManualMock) Handle(handler ReqHandlerManualMock) {
	handler.Handle(storage.Reader, &storage.Writer)
}

// Mock aliases.
type ()

func TestCrudControllerCrudCommandsRoute(t *testing.T) {
	// repos
	recRepo := repository.RecordRepo[string]{}
	ttlRepo := repository3.ExpiryRecRepo[time.Time]{}
	rLimRepo := repository2.RateLimitRepo[string]{}

	// ports
	reqPort := ReqHandlerPortManualMock{}

	// services
	reqService := reqserv.New[CrudCmdProducerManualMock, *WriterManualMock](reqPort)
	cmdEx := cmdserv.New(&recRepo, &ttlRepo)
	rLim := rlimserv.New(&rLimRepo)
	ttl := ttlserv.New(&ttlRepo)

	// routes
	path := "/api/rate_limit"
	rateLimiterRoute := CrudCommandsRoute[
		string, CrudCmdProducerManualMock, string, *WriterManualMock,
	](path)

	// controllers
	controller :=
		New[CrudCmdProducerManualMock, *WriterManualMock](
			reqService, cmdEx, rLim, ttl,
		)

	// background routes execution
	controller.RunConfig(rateLimiterRoute)

	// asserts
	expectedSavedLimit, _ := rLimRepo.Get(user)
	assert.Equal(t, expectedSavedLimit, rateLimit)

	expectedRequestRes := result
	assert.Equal(t, expectedRequestRes, storage.Writer.data[0])
}
