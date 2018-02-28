package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const yamlFile = ""

func TestCreateConfig(t *testing.T) {
	c := CreateConfig()

	assert.NotNil(t, c)
}

func TestCreateFromYAML(t *testing.T) {
	c, err := CreateFromYAML(yamlFile)

	assert := assert.New(t)
	assert.NotNil(c)
	assert.Nil(err)
}
