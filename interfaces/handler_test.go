package interfaces

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
)

func TestGetSentMessage(t *testing.T) {
	handler := Routes()

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/sent").
		Expect().
		Status(http.StatusOK).JSON()
}

