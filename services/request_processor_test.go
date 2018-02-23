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
        // valid request should return valid response
	// with empty error
	assert := assert.New(t)
	assert.NotNil(resp1)

	// invalid request should return nil response
	// with error
	assert.Nil(err1)
	assert.NotNil(resp2)
	assert.NotNil(err2)
}
