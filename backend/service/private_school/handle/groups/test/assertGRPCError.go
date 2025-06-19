package groups_handle_test

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func AssertGRPCError(t *testing.T, err error, expectedCode codes.Code, expectedMessage string) {
	t.Helper()
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	st, ok := status.FromError(err)
	if !ok {
		t.Fatalf("expected gRPC error, got: %T", err)
	}

	assert.Equal(t, expectedCode, st.Code(), "unexpected status code")
	assert.Equal(t, expectedMessage, st.Message(), "error message mismatch")
}
