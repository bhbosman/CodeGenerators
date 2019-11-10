package scopedObjects

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"strings"
)

type TypeDeclarator struct {
	TypeSpecBase
	simpleTypeSpec ScopingInterfaces.IDeclaredType
}

func NewTypeDeclarator(simpleTypeSpec ScopingInterfaces.IDeclaredType, declarator ScopingInterfaces.IDeclarator) (*TypeDeclarator, error) {
	sl := make([]string, 0, 0)
	for decl := declarator; decl != nil; decl = decl.GetNext() {
		sl = append(sl, decl.GetIdentifier())
	}
	id := strings.Join(sl, ",")
	return &TypeDeclarator{
		TypeSpecBase: NewTypeSpecBase(
			simpleTypeSpec,
			nil,
			id,
			ScopingInterfaces.TypeDeclaratorIdlType,
			false,
			false,
			false,
			false),
		simpleTypeSpec: simpleTypeSpec,
	}, nil
}

func (self *TypeDeclarator) TypeSpec() ScopingInterfaces.IDeclaredType {
	return self.simpleTypeSpec
}
func (self *TypeDeclarator) String() string {
	return fmt.Sprintf("TypeDeclarator: %v, %v", self.TypeSpecBase.String(), self.Identifier)
}
