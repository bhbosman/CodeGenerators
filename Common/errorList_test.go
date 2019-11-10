package Common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorList(t *testing.T) {
	t.Run("Test 0", func(t *testing.T) {
		l := newErrorList()
		assert.Len(t, l.list, 0)
		assert.NoError(t, l.Error())
	})

	t.Run("Test 1", func(t *testing.T) {
		l := newErrorList()
		l.Add(fmt.Errorf("some error"))
		assert.Len(t, l.list, 1)
		assert.Error(t, l.Error())
	})
}
