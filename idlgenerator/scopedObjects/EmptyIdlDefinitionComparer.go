package scopedObjects

import "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"

type EmptyIdlDefinitionComparer struct {
}

func (self *EmptyIdlDefinitionComparer) Compare(x, y ScopingInterfaces.IIdlDefinition) (ScopingInterfaces.IIdlDefinition, error) {
	return y, nil
}
