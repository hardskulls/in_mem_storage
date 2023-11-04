package crud_controller_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	cmdserv "in_mem_storage/internal/application/service/crud_cmd_executor"
	logger2 "in_mem_storage/internal/application/service/logger"
	rlimserv "in_mem_storage/internal/application/service/rate_limiter"
	reqserv "in_mem_storage/internal/application/service/server"
	ttlserv "in_mem_storage/internal/application/service/time_to_live"
	reqhdnl "in_mem_storage/internal/domain/incoming_request/value_object"
	"in_mem_storage/internal/domain/log/value_object/log_record"
	rlim "in_mem_storage/internal/domain/rate_limiter/value_object"
	"in_mem_storage/internal/infrastructure/db/in_mem/built_in/sync_map/service/command_executor/repository"
	repository2 "in_mem_storage/internal/infrastructure/db/in_mem/built_in/sync_map/service/rate_limiter/repository"
	repository3 "in_mem_storage/internal/infrastructure/db/in_mem/built_in/sync_map/service/time_to_live/repository"
	"in_mem_storage/internal/presentation/controllers/net/http/crud_controller"
	"sync"
	"testing"
	"time"
)

var storage = ReqRespStorageManualMock{
	Reader: ReaderManualMock{},
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
type ReaderManualMock struct{}

func (r ReaderManualMock) ProduceRateLim() (rlim.RateLimit, error) {
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
	Reader ReaderManualMock
	Writer WriterManualMock
}

// Request handler.
type ReqHandlerManualMock = reqhdnl.ReqHandler[ReaderManualMock, *WriterManualMock]

type ReqHandlerPortManualMock struct{}

func (r ReqHandlerPortManualMock) Handle(handler ReqHandlerManualMock) {
	handler.Handle(storage.Reader, &storage.Writer)
}

// Log record.
type LogRecordManualMock struct{}

func (l LogRecordManualMock) LogRecord(_ log_record.DefaultLogRecord) {}

func TestCrudControllerRateLimitRoute(t *testing.T) {
	recRepo := repository.RecordRepo[string]{}
	ttlRepo := repository3.ExpiryRecRepo[time.Time]{}
	rLimRepo := repository2.RateLimitRepo[string]{}

	reqAdapter := ReqHandlerPortManualMock{}
	logRecAdapter := LogRecordManualMock{}

	reqServ := reqserv.New[ReaderManualMock, *WriterManualMock](reqAdapter)
	cmdExServ := cmdserv.New(&recRepo, &ttlRepo)
	rLimServ := rlimserv.New(&rLimRepo)
	ttlServ := ttlserv.New(&ttlRepo)
	logServ := logger2.New(&logRecAdapter)

	path := "/api/rate_limit"
	rateLimiterRoute := crud_controller.RateLimiterRoute[
		ReaderManualMock, string, *WriterManualMock,
	](path)
	controller := crud_controller.New[ReaderManualMock, *WriterManualMock](
		reqServ, cmdExServ, rLimServ, ttlServ, logServ,
	)

	controller.RunConfig(rateLimiterRoute)

	expectedSavedLimit, _ := rLimRepo.Get(user)
	assert.Equal(t, expectedSavedLimit, rateLimit)

	expectedRequestRes := result
	assert.Equal(t, expectedRequestRes, storage.Writer.data[0])
}
