package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const yamlFile = ""

// test constructor
func TestCreateConfig(t *testing.T) {
	c := CreateConfig()

	assert.NotNil(t, c)
}

// create config from YAML
func TestCreateFromYAML(t *testing.T) {
	c, err := CreateFromYAML(yamlFile)

	assert := assert.New(t)
	assert.NotNil(c)
	assert.Nil(err)
}
