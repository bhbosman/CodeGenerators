package Extensions

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
)

type typeValueHelper struct {
}

var TypeValueHelper typeValueHelper

func (self typeValueHelper) TypeValueForDefinedType(DefinedType interfaces.IDefinedType) string {
	value := func() string {
		_, packageName, typeName := DefinedType.GetPackageName()
		if packageName != "" {
			return fmt.Sprintf("%v.%v", packageName, typeName)
		}
		return DefinedType.GetName()
	}()
	return func(value string) string {
		if DefinedType.Kind() == interfaces.Struct {
			return "*" + value
		}
		return value
	}(value)
}
