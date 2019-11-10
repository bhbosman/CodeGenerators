package scoping

import (
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"
)

type DefaultTypeService struct {
	declaredTypes map[string]ScopingInterfaces.IDeclaredType
}

func NewDefaultTypeService() *DefaultTypeService {
	declaredTypes := make(map[string]ScopingInterfaces.IDeclaredType)
	for i := ScopingInterfaces.PrimitiveTypesBegin; i <= ScopingInterfaces.PrimitiveTypesEnd; i++ {
		s01 := i.String()
		if s01 != "" {
			declaredTypes[s01] = scopedObjects.NewDeclaredType(i, "(built-in)", 0, 0)
		}
		s02 := i.IDLToken()
		if s02 != "" {
			declaredTypes[s02] = scopedObjects.NewDeclaredType(i, "(built-in)", 0, 0)
		}
	}
	return &DefaultTypeService{
		declaredTypes: declaredTypes,
	}
}

func (self *DefaultTypeService) Find(s string) ScopingInterfaces.IDeclaredType {
	dt, ok := self.FindOk(s)
	if ok {
		return dt
	}
	return nil
}

func (self *DefaultTypeService) FindOk(s string) (dt ScopingInterfaces.IDeclaredType, ok bool) {
	dt, ok = self.declaredTypes[s]
	return dt, ok
}

func (self *DefaultTypeService) Iterate(cb func(key string, declaredType ScopingInterfaces.IDeclaredType)) {
	if cb != nil {
		for key, value := range self.declaredTypes {
			cb(key, value)
		}
	}
}
