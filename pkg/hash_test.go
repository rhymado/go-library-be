package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pwd = "12345"
var hashed string
var hc = InitHashConfig().UseDefaultConfig()

func TestGenHashedPassword(t *testing.T) {
	hp, err := hc.GenHashedPassword(pwd)
	assert.NoError(t, err, "Error During Hashing")
	assert.NotEqual(t, pwd, hp, "Failed Hash Process")
	hashed = hp
}

func TestComparePasswordAndHash(t *testing.T) {
	t.Run("Case Password Valid", func(t *testing.T) {
		valid, err := hc.ComparePasswordAndHash(pwd, hashed)
		assert.NoError(t, err, "Error during Comparation")
		assert.True(t, valid, "Failed Comparation")
	})
	t.Run("Case Password Invalid", func(t *testing.T) {
		anotherPwd := "54321"
		valid, err := hc.ComparePasswordAndHash(anotherPwd, hashed)
		assert.NoError(t, err, "Error during Comparation")
		assert.False(t, valid, "Failed Comparation")
	})
}
