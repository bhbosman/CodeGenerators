package Common

import (
	"fmt"
	"strings"
)

// errorList keeps a list of errors in a string array. From an instance of errorList, an error will be created if there
// is an entry in the array.
// This may be extended to keep more information re the error, such as the call stack. There is some good example on the web
type errorList struct {
	list []string
}

// Error: Create an error, if there is an entry in the array
func (self *errorList) Error() error {
	if len(self.list) > 0 {
		return fmt.Errorf(strings.Join(self.list, "\n"))
	}
	return nil
}

// Add: Add an error to the list
func (self *errorList) Add(err error) {
	if err != nil {
		self.list = append(self.list, err.Error())
	}
	//return err
}

// An internal constructor used to create this error list
func newErrorList() *errorList {
	return &errorList{
		list: make([]string, 0, 16),
	}
}
