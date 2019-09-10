package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestM(t *testing.T) {
	tempHost := Host
	tempListenAndServe := listenAndServe

	listenAndServe = func(addr string, handler http.Handler) error {
		panic("Failed")
	}
	assert.PanicsWithValuef(t, "Failed", main, "Expected %v got %v", "Failed", main)

	Host = "invalid"
	assert.NotPanicsf(t, main, "Expected %v got %v", nil, main)
	Host = tempHost
	listenAndServe = tempListenAndServe
}
