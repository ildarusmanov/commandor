package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRequestProcessor(t *testing.T) {
	p := CreateRequestProcessor(validator, executor)

	assert.NotNil(t, p)
}
