package crud_controller_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	cmdserv "in_mem_storage/application/service/crud_cmd_executor"
	logger2 "in_mem_storage/application/service/logger"
	rlimserv "in_mem_storage/application/service/rate_limiter"
	reqserv "in_mem_storage/application/service/server"
	ttlserv "in_mem_storage/application/service/time_to_live"
	reqhdnl "in_mem_storage/domain/incoming_request/value_object"
	rlim "in_mem_storage/domain/rate_limiter/value_object"
	"in_mem_storage/infrastructure/db/in_mem/built_in/sunc_map/service/command_executor/repository"
	repository2 "in_mem_storage/infrastructure/db/in_mem/built_in/sunc_map/service/rate_limiter/repository"
	repository3 "in_mem_storage/infrastructure/db/in_mem/built_in/sunc_map/service/time_to_live/repository"
	crudctrl "in_mem_storage/presentation/controllers/net/http/crud_controller"
	"sync"
	"testing"
	"time"
	"in_mem_storage/infrastructure/io/system/built_in/console/service/logger/adapter"
)

var storage = ReqRespStorageManualMock{
	Reader: RateLimitProducerManualMock{},
	Writer: WriterManualMock{data: make([]string, 0)},
}
var user = "user_12345"
var rateLimit = rlim.RateLimit{
	For:      user,
	LastUsed: time.Now(),
	Limit:    time.Millisecond * 500,
}
var result = fmt.Sprintf("[RateLimitOperationSuccess] Your rate limit is, %v", rateLimit)

// Request.
type RateLimitProducerManualMock struct{}

func (r RateLimitProducerManualMock) ProduceRateLim() (rlim.RateLimit, error) {
	return rateLimit, nil
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
	Reader RateLimitProducerManualMock
	Writer WriterManualMock
}

// Request handler.
type ReqHandlerManualMock = reqhdnl.ReqHandler[RateLimitProducerManualMock, *WriterManualMock]

type ReqHandlerPortManualMock struct{}

func (r ReqHandlerPortManualMock) Handle(handler ReqHandlerManualMock) {
	handler.Handle(storage.Reader, &storage.Writer)
}

func TestCrudControllerRateLimitRoute(t *testing.T) {
	// repos
	recRepo := repository.RecordRepo[string]{}
	ttlRepo := repository3.ExpiryRecRepo[time.Time]{}
	rLimRepo := repository2.RateLimitRepo[string]{}

	//  adapters
	reqAdapter := ReqHandlerPortManualMock{}
	logRecAdapter := adapter.LogRecordAdapter{}

	// services
	reqService := reqserv.New[RateLimitProducerManualMock, *WriterManualMock](reqAdapter)
	cmdEx := cmdserv.New(&recRepo, &ttlRepo)
	rLim := rlimserv.New(&rLimRepo)
	ttl := ttlserv.New(&ttlRepo)
	logger := logger2.New(&loggerRepo)

	// routes
	path := "/api/rate_limit"
	rateLimiterRoute := crudctrl.NewRateLimiterRoute[RateLimitProducerManualMock, string, *WriterManualMock](path)

	// controllers
	controller := crudctrl.
		New[RateLimitProducerManualMock, *WriterManualMock](
		reqService, cmdEx, rLim, ttl, logger,
	)

	// background routes execution
	controller.RunConfig(rateLimiterRoute)

	// asserts
	expectedSavedLimit, _ := rLimRepo.Get(user)
	assert.Equal(t, expectedSavedLimit, rateLimit)

	expectedRequestRes := result
	assert.Equal(t, expectedRequestRes, storage.Writer.data[0])
}
