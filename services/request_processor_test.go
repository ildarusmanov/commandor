package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// test constructor
func TestCreateRequestProcessor(t *testing.T) {
	p := CreateRequestProcessor(validator, executor)

	assert.NotNil(t, p)
}

// test request processing
func TestPequestProcessorProcess(t *testing.T) {
	p := CreateRequestProcessor(validator, executor)

	resp1, err1 := p.Process(validRequest)
	resp2, err2 := p.Process(invalidRequest)

	assert := assert.New(t)
	assert.NotNil(resp1)
	assert.Nil(err1)
	assert.NotNil(resp2)
	assert.NotNil(err2)
}
