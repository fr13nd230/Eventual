package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomID(t *testing.T) {
	for range 1000 {
		r := generateUniqueID()
		assert.IsType(t, "", r)
        assert.Greater(t, len(r), 0)
    }
	t.Parallel()
}
