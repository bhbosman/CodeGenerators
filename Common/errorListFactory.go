package Common

import "fmt"

var ErrorListFactory errorListFactory

type errorListFactory struct {
	other func() IErrorList
}

func (self *errorListFactory) NewErrorListFunc(cb func(errorList IErrorList)) error {
	if cb == nil {
		return fmt.Errorf("no callback assigned")
	}
	list := func() IErrorList {
		if self.other != nil {
			return self.other()
		}
		return newErrorList()
	}()
	if list == nil {
		return fmt.Errorf("list generated was nil")
	}
	cb(list)
	return list.Error()
}

func (self *errorListFactory) Replace(other func() IErrorList) {
	self.other = other
}

func (self *errorListFactory) Reset() {
	self.other = nil
}
