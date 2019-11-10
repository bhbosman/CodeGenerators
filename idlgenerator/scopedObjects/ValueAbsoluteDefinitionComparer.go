package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type ValueAbsoluteDefinitionComparer struct {
}

func (self *ValueAbsoluteDefinitionComparer) Compare(x, y ScopingInterfaces.IIdlDefinition) (ScopingInterfaces.IIdlDefinition, error) {
	xx, ok := x.(ScopingInterfaces.IValueAbsoluteDefinition)
	if !ok {
		return nil, fmt.Errorf("type not a IInterfaceDcl")
	}

	yy, ok := y.(ScopingInterfaces.IValueAbsoluteDefinition)
	if !ok {
		return nil, fmt.Errorf("type not a IInterfaceDcl")
	}

	return func(x, y ScopingInterfaces.IValueAbsoluteDefinition) (ScopingInterfaces.IIdlDefinition, error) {
		if x.Forward() {
			b := true
			b = b && x.GetName() == y.GetName()
			if y.Forward() {
				// double forward declaration
				return x, nil
			} else {
				return y, nil
			}
		}
		// self.forward is false
		if y.Forward() {
			b := true
			b = b && x.GetName() == y.GetName()
			if y.Forward() {
				// later forward declaration
				return y, nil
			}
		}
		// self.forward is false && other.Forward() is false
		return nil, fmt.Errorf("double declaration of %v", x.GetName())
	}(xx, yy)
}