package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type IdlPlaceHolder struct {
	Number  int
	idlName string
	idlType ScopingInterfaces.IIdlDefinition
}

func (self *IdlPlaceHolder) DoCompare(x, idlDefinition ScopingInterfaces.IIdlDefinition) (ScopingInterfaces.IIdlDefinition, error) {
	factory, ok := x.(ScopingInterfaces.IIdlCompareFactory)
	if !ok {
		//return nil
		return nil, fmt.Errorf("no comparer found for %v. (%v)", self.idlName, self.idlType)
	}

	instance := factory.Create()
	definition, err := instance.Compare(x, idlDefinition)
	if err != nil {
		return nil, err
	}
	if definition != nil {
		self.idlType = definition
	}
	return self.idlType, nil
}

func NewIdlPlaceHolder(number int, idlName string, idlType ScopingInterfaces.IIdlDefinition) *IdlPlaceHolder {
	return &IdlPlaceHolder{
		Number:  number,
		idlName: idlName,
		idlType: idlType,
	}
}

func (self *IdlPlaceHolder) IdlName() string {
	return self.idlName
}

func (self *IdlPlaceHolder) IdlType() ScopingInterfaces.IIdlDefinition {
	return self.idlType
}
