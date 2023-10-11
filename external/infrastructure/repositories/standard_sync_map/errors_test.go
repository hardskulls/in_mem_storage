package standard_sync_map

import (
	"strings"
	"testing"
)

func TestDeleteError(t *testing.T) {
	err := DeleteError().OccuredAt()
	// [!!] Line number is crutial here [!!]
	hasValidFileLineInfo := strings.HasSuffix(err, "errors_test.go:9")
	if !hasValidFileLineInfo {
		t.Fatalf("expected \"\", found  \"%v\"", err)
	}
}

func TestDGetError(t *testing.T) {
	err := GetError().OccuredAt()
	// [!!] Line number is crutial here [!!]
	hasValidFileLineInfo := strings.HasSuffix(err, "errors_test.go:17")
	if !hasValidFileLineInfo {
		t.Fatalf("expected \"\", found  \"%v\"", err)
	}
}
