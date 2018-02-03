package storage

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestCreateMemoryStorage(t *testing.T) {
  s, err := CreateMemoryStorage(options)

  assert := assert.New(t)
  assert.Nil(err)
  assert.NotNil(s)
}

func TestMemoryStorageSaveCommand(t *testing.T) {
  t.Error("Test not implemented")
}

func TestMemoryStorageFindCommand(t *testing.T) {
  t.Error("Test not implemented")
}

