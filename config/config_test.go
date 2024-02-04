package config_test

import (
	"github.com/mbn18/scan/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	conf := config.Get()

	str := conf.GetConnString()
	assert.Equal(t, "postgres://cypago_user:pass@10.0.0.16/cypago?sslmode=disable", str)
}
