package services

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

const token = "token"

func TestCreateRequestValidator(t *testing.T) {
    v := CreateRequestValidator(token)

    assert.NotNil(t, v)
}

func TestRequestValidatorCreateSignature(t *testing.T) {
    v := CreateRequestValidator(token)
    s := v.CreateSignature(timestamp)

    assert.NotNil(t, s)
}

func TestRequestValidatorValidateSignature(t *testing.T) {
    v := CreateRequestValidator(token)
    validSignature := v.CreateSignature(timestamp)
    invalidSignature := validSignature + "1"

    assert := assert.New(t)
    assert.True(v.ValidateSignature(validSignature))
    assert.False(v.ValidateSignature(invalidSignature))
}
