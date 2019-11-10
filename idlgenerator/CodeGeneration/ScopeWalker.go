package CodeGeneration

import (
	"fmt"
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"go.uber.org/multierr"
	"strings"
	"unicode"
)

type IFileWriter interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type GolangCodeGenerator struct {
	Logger IFileWriter
	indent int
}

func NewGenerateCodeGolang(logger IFileWriter) *GolangCodeGenerator {
	return &GolangCodeGenerator{
		Logger: logger,
	}
}
func (self GolangCodeGenerator) Generate(dcl si.ITypeSpec) error {
	var err error
	for typeSpec := dcl; typeSpec != nil; typeSpec, _ = typeSpec.GetNextTypeSpec() {
		err = multierr.Append(err, self.InternalGenerate(typeSpec))
	}
	return err

}

func (self GolangCodeGenerator) InternalGenerate(dcl si.ITypeSpec) error {
	self.BlankLine()
	self.WriteComment("line %v:%v", dcl.GetFileName(), dcl.GetRow())
	switch dcl.GetKind() {

	case si.Attr_specIdlType:
		attributeDecl, ok := dcl.(si.IAttributeDcl)
		if ok {
			return self.GenerateAttributeDcl(attributeDecl)
		}

	case si.ExceptionIdlType:
		exceptionDecl, ok := dcl.(si.IIdlException)
		if ok {
			return self.GenerateExceptionDcl(exceptionDecl)
		}

	case si.ConstDclType:
		constantDecl, ok := dcl.(si.IIdlConstDcl)
		if ok {
			return self.GenerateConstantDcl(constantDecl)
		}
	case si.ModuleIdlType:
		moduleDcl, ok := dcl.(si.IIdlModuleDcl)
		if ok {
			return self.GenerateModuleDcl(moduleDcl)
		}
	case si.StructIdlType:
		structType, ok := dcl.(si.IStructType)
		if ok {
			return self.GenerateStructDcl(structType)
		}
	case si.RWEnumIdlType:
		enumType, ok := dcl.(si.IEnumType)
		if ok {
			return self.GenerateEnumDcl(enumType)
		}

	case si.InterfaceIdlType:
		interfaceDcl, ok := dcl.(si.IInterfaceDcl)
		if ok {
			return self.GenerateInterfaceDcl(interfaceDcl)
		}
	case si.Op_dclIdlType:
		operation, ok := dcl.(si.IOperationDeclarations)
		if ok {
			return self.GenerateOperationDcl(operation)
		}
	case si.IdlValue_Abs_DefType:
		value, ok := dcl.(si.IValueAbsoluteDefinition)
		if ok {
			return self.GenerateValueAbsolute(value)
		}

	case si.TypeDeclaratorIdlType:
		typeDecl, ok := dcl.(si.ITypeDeclarator)
		if ok {
			return self.GenerateTypeDcl(typeDecl)
		}
	default:
		self.WriteLine("// No declare Type")
	}
	return nil
}
func (self GolangCodeGenerator) GenerateModuleDcl(dcl si.IIdlModuleDcl) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", dcl))
	var err error

	err = multierr.Append(
		err,
		dcl.Iterate(func(typeSpec si.ITypeSpec) error {
			return self.InternalGenerate(typeSpec)
		}))

	return err

}
func (self GolangCodeGenerator) WriteLine(f string, p ...interface{}) {
	tabs := strings.Repeat("\t", self.indent)
	s := fmt.Sprintf(f, p...)
	self.Logger.Printf("%v%v\n", tabs, s)
}

func (self GolangCodeGenerator) WriteComment(f string, p ...interface{}) {
	tabs := strings.Repeat("\t", self.indent)
	s := fmt.Sprintf(f, p...)
	self.Logger.Printf("// %v%v\n", tabs, s)
}

func (self GolangCodeGenerator) BlankLine() {
	self.Logger.Printf("\n")
}

func (self *GolangCodeGenerator) IncIndent() {
	self.indent++
}

func (self *GolangCodeGenerator) DecIndent() {
	self.indent--
}
func (self GolangCodeGenerator) exportDclName(dcl si.IDeclaredType) string {
	s := strings.Split(dcl.GetName(), "::")
	return strings.Join(s, "_")
}
func (self GolangCodeGenerator) exportMemberName(name string) string {
	b := []byte(name)
	b[0] = byte(unicode.ToUpper(rune(name[0])))
	return string(b)
}

func (self GolangCodeGenerator) GenerateStructDcl(dcl si.IStructType) error {
	var err error
	exportName := self.exportDclName(dcl)
	self.WriteComment("Struct Decl: %v", exportName)
	self.WriteLine("type %v struct {", exportName)
	self.IncIndent()
	if dcl.Members() != nil {
		for _, memberInfo := range dcl.Members().GetMembers() {
			self.WriteLine("%v %v", self.exportMemberName(memberInfo.GetId()), self.exportDclName(memberInfo.GetTypeSpec()))
		}
	}
	self.DecIndent()
	self.WriteLine("}")

	return err
}
func (self GolangCodeGenerator) GenerateEnumDcl(enumType si.IEnumType) error {

	for m := enumType.Enumerator(); m != nil; m = m.Next() {

	}
	return nil
}
func (self GolangCodeGenerator) buildDeclarationName(scope, name string) string {
	if scope == "" {
		return name
	}
	return fmt.Sprintf("%v::%v", scope, name)

}
func (self GolangCodeGenerator) GenerateInterfaceDcl(dcl si.IInterfaceDcl) error {

	var err error

	return err
}

func (self GolangCodeGenerator) GenerateOperationDcl(dcl si.IOperationDeclarations) error {
	var err error
	self.Logger.Println(fmt.Sprintf("(%d): %v", dcl))

	for param := dcl.GetParams(); param != nil; param = param.GetNextParameterDeclarations() {
		self.Logger.Println(fmt.Sprintf("(%d): %v", param))
	}
	return err
}

func (self GolangCodeGenerator) GenerateTypeDcl(dcl si.ITypeDeclarator) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", dcl))
	return nil
}

func (self GolangCodeGenerator) GenerateValueAbsolute(definition si.IValueAbsoluteDefinition) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", definition))
	return definition.Iterate(func(typeSpec si.ITypeSpec) error {
		return self.InternalGenerate(typeSpec)
	})
}

func (self GolangCodeGenerator) GenerateConstantDcl(constDcl si.IIdlConstDcl) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", constDcl))
	return nil
}

func (self GolangCodeGenerator) GenerateExceptionDcl(exceptionType si.IIdlException) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", exceptionType))

	members := exceptionType.GetMembers()
	if members != nil {
		for _, memberInformation := range members.GetMembers() {
			self.Logger.Println(fmt.Sprintf("(%d): %v", memberInformation))
		}
	}
	return nil
}

func (self GolangCodeGenerator) GenerateAttributeDcl(attributeDcl si.IAttributeDcl) error {
	self.Logger.Println(fmt.Sprintf("(%d): %v", attributeDcl))
	return nil
}
