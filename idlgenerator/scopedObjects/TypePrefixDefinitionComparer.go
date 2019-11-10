package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type TypePrefixDefinitionComparer struct {
}

func (self *TypePrefixDefinitionComparer) Compare(x, y ScopingInterfaces.IIdlDefinition) (ScopingInterfaces.IIdlDefinition, error) {
	xx, ok := x.(ScopingInterfaces.ITypePrefixDefinition)
	if !ok {
		return nil, fmt.Errorf("type not a ITypePrefixDefinition")
	}

	yy, ok := y.(ScopingInterfaces.ITypePrefixDefinition)
	if !ok {
		return nil, fmt.Errorf("type not a ITypePrefixDefinition")
	}
	return func(xx, yy ScopingInterfaces.ITypePrefixDefinition) (ScopingInterfaces.IIdlDefinition, error) {
		b := true
		b = b && xx.GetName() == yy.GetName()
		b = b && xx.StringLiteral() == yy.StringLiteral()
		if !b {
			return nil, fmt.Errorf("object declared more than once and they are not the same")
		}
		return nil, nil

	}(xx, yy)
}
