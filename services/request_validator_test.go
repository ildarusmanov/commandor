package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// token example
const token = "token"

// test constructor
func TestCreateRequestValidator(t *testing.T) {
	v := CreateRequestValidator(token)

	assert.NotNil(t, v)
}

// test token create method
func TestRequestValidatorCreateSignature(t *testing.T) {
	v := CreateRequestValidator(token)
	s := v.CreateSignature(timestamp)
        // new token should not be empty
	assert.NotNil(t, s)
}

// test validate method
func TestRequestValidatorValidateSignature(t *testing.T) {
	v := CreateRequestValidator(token)
	validSignature := v.CreateSignature(timestamp)
	invalidSignature := validSignature + "1"

	assert := assert.New(t)
	// valid token should finish with True
	assert.True(v.ValidateSignature(validSignature))
        // invali token should finish with False
	assert.False(v.ValidateSignature(invalidSignature))
}
