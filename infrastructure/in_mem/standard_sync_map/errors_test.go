package standard_sync_map

import (
	"strings"
	"testing"
)

func TestDeleteError(t *testing.T) {
	err := DeleteError().OccurredAt()
	// [!!] Line number is crucial here [!!]
	hasValidFileLineInfo := strings.HasSuffix(err, "errors_test.go:9")
	if !hasValidFileLineInfo {
		t.Fatalf("expected \"\", found  \"%v\"", err)
	}
}

func TestDGetError(t *testing.T) {
	err := GetError().OccurredAt()
	// [!!] Line number is crucial here [!!]
	hasValidFileLineInfo := strings.HasSuffix(err, "errors_test.go:17")
	if !hasValidFileLineInfo {
		t.Fatalf("expected \"\", found  \"%v\"", err)
	}
}

func comp[C comparable](l C, r C) bool {
	return l == r
}

func foo() {
	s := "foo"
	res := comp(s, "foo")
	println(res)
}
