package CustomError

import (
	"errors"
	"testing"
)

func TestError(t *testing.T) {
	err := NewError("test error")
	var cerr *CError
	errors.As(err, &cerr)
	if cerr.Line != 9 {
		t.Errorf("expected line 9, got %d", cerr.Line)
	}
	if cerr.Message != "test error" {
		t.Errorf("expected test error, got %s", cerr.Message)
	}
}
