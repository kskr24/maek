package auth_test

import (
	"context"
	"testing"

	approvals "github.com/customerio/go-approval-tests"

	"github.com/karngyan/maek/zarf/tests"
	"github.com/stretchr/testify/assert"

	"github.com/karngyan/maek/domains/auth"
)

func TestLogin(t *testing.T) {
	user, session, err := auth.CreateDefaultAccountWithUser(context.Background(), "Karn", "karn@maek.ai", "test-password", "1.2.3.4", "Mozilla/5.0")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotNil(t, session)

	rr, err := tests.Post("/v1/auth/login", map[string]any{
		"email":    "karn@maek.ai",
		"password": "test-password",
	})
	assert.Nil(t, err)
	assert.Equal(t, 200, rr.Code)

	approvals.VerifyJSONBytes(t, rr.Body.Bytes())
	assert.Contains(t, rr.Header().Get("Set-Cookie"), "HttpOnly; Secure; SameSite=Strict")
}