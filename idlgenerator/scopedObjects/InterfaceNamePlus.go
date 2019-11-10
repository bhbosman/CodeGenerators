package scopedObjects

import "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"

type InterfaceNamePlus struct {
}

func (self *InterfaceNamePlus) Next(identifier string) ScopingInterfaces.IInterfaceNamePlus {
	return self
}

func NewInterfaceNamePlus(identifier string) (*InterfaceNamePlus, error) {
	return &InterfaceNamePlus{}, nil
}
