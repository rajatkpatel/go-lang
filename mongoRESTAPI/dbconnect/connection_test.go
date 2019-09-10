package dbconnect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var createSessionInput = []struct {
	host  string
	isErr bool
}{
	{"mongodb://127.0.0.1:27017", false},
	{"invalid", true},
}

func TestCreateSession(t *testing.T) {
	for _, item := range createSessionInput {
		_, err := CreateSession(item.host)
		assert.Equalf(t, item.isErr, err != nil, "Expected %v got %v", item.isErr, err != nil)
	}

}
