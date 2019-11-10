package Common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorListFactory(t *testing.T) {
	ErrorListFactory.Reset()
	assert.Nil(t, ErrorListFactory.other)

	t.Run("Default create with no callback", func(t *testing.T) {
		err := ErrorListFactory.NewErrorListFunc(nil)
		assert.NotNil(t, err)
	})

	t.Run("Default create with callback and no error", func(t *testing.T) {
		err := ErrorListFactory.NewErrorListFunc(func(errorList IErrorList) {

		})
		assert.Nil(t, err)
	})

	t.Run("Default create with callback with error", func(t *testing.T) {
		err := ErrorListFactory.NewErrorListFunc(func(errorList IErrorList) {
			errorList.Add(fmt.Errorf("some error"))
		})
		assert.NotNil(t, err)
	})

	t.Run("Use other implementation of IErrorList", func(t *testing.T) {
		ErrorListFactory.Replace(func() IErrorList {
			return nil
		})
		defer func() {
			ErrorListFactory.Reset()
		}()

		err := ErrorListFactory.NewErrorListFunc(func(errorList IErrorList) {
			errorList.Add(fmt.Errorf("some error"))
		})
		assert.NotNil(t, err)
	})

	t.Run("Use other implementation of IErrorList", func(t *testing.T) {
		mockList := newMockErrorList()
		ErrorListFactory.Replace(func() IErrorList {
			return mockList
		})
		defer func() {
			ErrorListFactory.Reset()
		}()

		err := ErrorListFactory.NewErrorListFunc(func(errorList IErrorList) {
			errorList.Add(fmt.Errorf("some error"))
		})
		assert.NotNil(t, err)
		assert.Len(t, mockList.l, 1)
	})
}

type mockErrorList struct {
	l []string
}

func newMockErrorList() *mockErrorList {
	return &mockErrorList{
		l: make([]string, 0, 0),
	}
}

func (self *mockErrorList) Add(err error) {
	self.l = append(self.l, err.Error())
}

func (self *mockErrorList) Error() error {
	if len(self.l) == 0 {
		return nil
	}
	return fmt.Errorf("some error with %v items", len(self.l))
}
