package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
)

type AttributeDcl struct {
	TypeSpecBase
	attrDeclarator ScopingInterfaces.IAttrDeclarator
	declaredType   ScopingInterfaces.IBaseDeclaredType
	readOnly       bool
}

func NewAttributeDcl(fileInformation ScopingInterfaces.IFileInformation, attrDeclarator ScopingInterfaces.IAttrDeclarator, declaredType ScopingInterfaces.IBaseDeclaredType, readOnly bool) (*AttributeDcl, error) {
	return &AttributeDcl{
		TypeSpecBase: NewTypeSpecBase(
			fileInformation,
			nil,
			attrDeclarator.GetName(),
			ScopingInterfaces.Attr_specIdlType,
			false,
			false,
			false,
			false),
		attrDeclarator: attrDeclarator,
		declaredType:   declaredType,
		readOnly:       readOnly,
	}, nil
}

func (self *AttributeDcl) DeclaredType() ScopingInterfaces.IBaseDeclaredType {
	return self.declaredType
}

func (self *AttributeDcl) AttrDeclarator() ScopingInterfaces.IAttrDeclarator {
	return self.attrDeclarator
}

func (self *AttributeDcl) ReadOnly() bool {
	return self.readOnly
}

func (self *AttributeDcl) String() string {
	return fmt.Sprintf("AttributeDcl: %v, %v: %v", self.TypeSpecBase.String(), self.Identifier, self.declaredType.GetName())
}
