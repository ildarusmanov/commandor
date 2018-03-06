package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	yamlFile      = "tests/fixtures/config.yml"
	yamlSecretKey = "test_secret_key"
	yamlTokenType = "tesst_token_type"
)

// test constructor
func TestCreateNewConfig(t *testing.T) {
	c := CreateNewConfig()

	assert.NotNil(t, c)
}

// create config from YAML
func TestCreateFromYAML(t *testing.T) {
	c, err := CreateFromYAML(yamlFile)

	assert := assert.New(t)

	assert.NotNil(c)
	assert.Nil(err)
	assert.Equal(config.GetSecretKey(), yamlSecretKey)
	assert.Equal(config.GetTokenType(), yamlTokenType)
}
