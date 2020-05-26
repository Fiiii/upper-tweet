package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_MainHealth(t *testing.T){
	testServer := httptest.NewServer(InitServer())
	defer testServer.Close()

	res, err := http.Get(fmt.Sprintf("%s/ping", testServer.URL))
	if err != nil {
		t.Fatalf("Error, %v", err)
	}

	assert.Equal(t, 200, res.StatusCode)
}