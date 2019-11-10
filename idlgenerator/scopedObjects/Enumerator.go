package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type Enumerator struct {
	FileInformationBase
	id   string
	next ScopingInterfaces.IEnumerator
}

func NewEnumerator(fileInformation ScopingInterfaces.IFileInformation, id string) (ScopingInterfaces.IEnumerator, error) {
	if id == "" {
		return nil, fmt.Errorf("invalid enum enumerator")
	}
	return &Enumerator{
		FileInformationBase: NewFileInformationBase02(fileInformation),
		id:                  id,
		next:                nil,
	}, nil
}

func (self *Enumerator) String() string {
	return fmt.Sprintf("Enumerator: %v, %v, ", self.FileInformationBase.String(), self.id)
}

func (self *Enumerator) Id() string {
	return self.id
}

func (self *Enumerator) SetLast(next ScopingInterfaces.IEnumerator) {
	self.next = next
}

func (self *Enumerator) Next() ScopingInterfaces.IEnumerator {
	return self.next
}

func (self *Enumerator) Last(last ScopingInterfaces.IEnumerator) ScopingInterfaces.IEnumerator {
	var iter ScopingInterfaces.IEnumerator = self
	for iter.Next() != nil {
		iter = iter.Next()
	}
	iter.SetLast(last)
	return self
}
