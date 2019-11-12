package CodeGeneration

import (
	"fmt"
	si "github.com/bhbosman/CodeGenerators/idlgenerator/ScopingInterfaces"
	"go.uber.org/multierr"
	"io"
	"os"
	"path"
	"strings"
	"unicode"
)

type IFileWriter interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

type GolangCodeGenerator struct {
	indent    int
	openFiles map[string]io.WriteCloser
}

func NewGenerateCodeGolang() *GolangCodeGenerator {
	return &GolangCodeGenerator{
		indent:    0,
		openFiles: make(map[string]io.WriteCloser),
	}
}

func (self GolangCodeGenerator) Generate(dcl si.ITypeSpec) error {
	var err error
	for typeSpec := dcl; typeSpec != nil; typeSpec, _ = typeSpec.GetNextTypeSpec() {
		err = multierr.Append(err, self.internalGenerate(nil, typeSpec))
	}
	return err
}

func (self *GolangCodeGenerator) Close() error {
	var err error = nil
	for _, v := range self.openFiles {
		err = multierr.Append(
			err,
			v.Close())
	}
	return err
}

func (self GolangCodeGenerator) findWriter(dcl si.ITypeSpec) (io.Writer, error) {
	idlFileName := dcl.GetFileName()
	if writer, ok := self.openFiles[idlFileName]; ok {
		return writer, nil
	}
	goFile := fmt.Sprintf("%v.%v", idlFileName, "go")
	file, err := os.Create(goFile)
	if err != nil {
		return nil, err
	}


	dir, _ := path.Split(idlFileName)
	splitFolders := strings.Split(dir, string(os.PathSeparator))

	self.writeLine(file, false, "package %v", splitFolders[len(splitFolders)-2])
	self.blankLine(file)
	self.writeComment(file,"Code generated by  DO NOT EDIT.")


	self.openFiles[idlFileName] = file
	return file, nil
}

func (self GolangCodeGenerator) internalGenerate(writer io.Writer, dcl si.ITypeSpec) error {
	switch dcl.GetKind() {
	case si.Attr_specIdlType:
		attributeDecl, ok := dcl.(si.IAttributeDcl)
		if ok {
			return self.generateAttributeDcl(writer, attributeDecl)
		}
	case si.ExceptionIdlType:
		exceptionDecl, ok := dcl.(si.IIdlException)
		if ok {
			newWriter, e := self.findWriter(dcl)
			if e != nil {
				
			}
			return self.generateExceptionDcl(newWriter, exceptionDecl)
		}
	case si.ConstDclType:
		constantDecl, ok := dcl.(si.IIdlConstDcl)
		if ok {
			return self.generateConstantDcl(writer, constantDecl)
		}
	case si.ModuleIdlType:
		moduleDcl, ok := dcl.(si.IIdlModuleDcl)
		if ok {
			return self.generateModuleDcl(writer, moduleDcl)
		}
	case si.StructIdlType:
		structType, ok := dcl.(si.IStructType)
		if ok {
			newWriter, e := self.findWriter(dcl)
			if e != nil {
				
			}
			return self.generateStructDcl(newWriter, structType)
		}
	case si.RWEnumIdlType:
		enumType, ok := dcl.(si.IEnumType)
		if ok {
			newWriter, e := self.findWriter(dcl)
			if e != nil {
				
			}
			return self.generateEnumDcl(newWriter, enumType)
		}

	case si.InterfaceIdlType:
		interfaceDcl, ok := dcl.(si.IInterfaceDcl)
		if ok {
			newWriter, e := self.findWriter(dcl)
			if e != nil {
				
			}
			return self.generateInterfaceDcl(newWriter, interfaceDcl)
		}
	case si.Op_dclIdlType:
		operation, ok := dcl.(si.IOperationDeclarations)
		if ok {
			return self.generateOperationDcl(writer, operation)
		}
	case si.IdlValue_Abs_DefType:
		value, ok := dcl.(si.IValueAbsoluteDefinition)
		if ok {
			newWriter, e := self.findWriter(dcl)
			if e != nil {
				
			}
			return self.generateValueAbsolute(newWriter, value)
		}

	case si.TypeDeclaratorIdlType:
		typeDecl, ok := dcl.(si.ITypeDeclarator)
		if ok {
			return self.generateTypeDcl(writer, typeDecl)
		}
	default:
		self.writeLineNumber(writer, dcl)
		self.writeLine(writer, true, "// No declare Type")
	}
	return nil
}

func (self GolangCodeGenerator) writeLineNumber(writer io.Writer, dcl si.ITypeSpec) {
	_, file := path.Split(dcl.GetFileName())
	self.writeComment(writer, "line %v:%v", file, dcl.GetRow())
}

func (self GolangCodeGenerator) generateModuleDcl(writer io.Writer, dcl si.IIdlModuleDcl) error {
	self.writeLineNumber(writer, dcl)
	var err error

	err = multierr.Append(
		err,
		dcl.Iterate(func(typeSpec si.ITypeSpec) error {
			return self.internalGenerate(writer, typeSpec)
		}))

	return err
}

func (self GolangCodeGenerator) exportDclName(dcl si.IDeclaredType) string {
	switch dcl.GetKind(){
	case si.RWfloatIdlType:
		return "float"
	case si.StringIdlType:
			return "string"
	case si.WideStringIdlType:
			return "string"
	case si.RWVoidType:
			return "void"
	case si.RWdoubleIdlType:
			return "float64"
	case si.RWlongRWdoubleIdlType:
			return "float64"
	case si.RWshortIdlType:
			return "int16"
	case si.RWlongIdlType:
			return "int32"
	case si.RWint8IdlType:
			return "int8"
	case si.RWlongRWlongIdlType:
			return "int64"
	case si.RWuint8IdlType:
			return "uint8"
	case si.RWunsignedRWshortIdlType:
			return "uint16"
	case si.RWunsignedRWlongIdlType:
			return "uint32"
	case si.RWunsignedRWlongRWlongIdlType:
			return "uint63"
	case si.RWcharIdlType:
			return "rune"
	case si.RWwcharIdlType:
			return "rune"
	case si.RWbooleanIdlType:
			return "bool"
	case si.RWoctetIdlType:
			return "int"
	case si.RWanyIdlType:
			return "CorbaAny"
	case si.RWObjectIdlType:
			return "CorbaObject"
	case si.RWint16IdlType:
			return "int16"
	case si.RWint32IdlType:
			return "int32"
	case si.RWint64IdlType:
			return "int64"
	case si.RWuint16IdlType:
			return "uint16"
	case si.RWuint32IdlType:
			return "uint32"
	case si.RWuint64IdlType:
			return "uint64"



	case si.RWEnumIdlType:
		s := strings.Split(dcl.GetName(), "::")
		return strings.Join(s, "_")

	case si.RWvaluebaseIdlType:
		s := strings.Split(dcl.GetName(), "::")
		return strings.Join(s, "_")


	default:
		s := strings.Split(dcl.GetName(), "::")
		return strings.Join(s, "_")
	}
}
func (self GolangCodeGenerator) exportMemberName(name string) string {
	b := []byte(name)
	b[0] = byte(unicode.ToUpper(rune(name[0])))
	return string(b)
}
func (self GolangCodeGenerator) incomingParamName(name string) string {
	b := []byte(name)
	b[0] = byte(unicode.ToLower(rune(name[0])))
	return string(b)
}

func (self GolangCodeGenerator) generateEnumDcl(writer io.Writer, dcl si.IEnumType) error {
	self.writeLineNumber(writer, dcl)
	for m := dcl.Enumerator(); m != nil; m = m.Next() {

	}
	return nil
}
func (self GolangCodeGenerator) buildDeclarationName(scope, name string) string {
	if scope == "" {
		return name
	}
	return fmt.Sprintf("%v::%v", scope, name)

}

func (self GolangCodeGenerator) generateValueAbsolute(writer io.Writer, dcl si.IValueAbsoluteDefinition) error {
	exportName := self.exportDclName(dcl)
	if dcl.Forward() {
		self.writeLineNumber(writer, dcl)
		self.writeComment(writer, "Forward ValueAbsolute Decl: %v", exportName)
		return nil
	}
	self.writeLineNumber(writer, dcl)
	self.writeComment(writer, "ValueAbsolute Decl: %v", exportName)
	self.writeLine(writer, true, "type %v interface {", exportName)
	self.incIndent(writer)
	err := multierr.Append(
		nil,
		dcl.Iterate(
			func(dcl si.ITypeSpec) error {
				return self.internalGenerate(writer, dcl)
			}))
	self.decIndent(writer)
	self.writeLine(writer, true, "}")
	return err

}

func (self GolangCodeGenerator) generateInterfaceDcl(writer io.Writer, dcl si.IInterfaceDcl) error {
	doBefore := func(dcl si.ITypeSpec) bool {
		switch dcl.GetKind() {
		case si.StructIdlType, si.ExceptionIdlType:
			return true
		default:
			return false
		}
	}
	err := dcl.Iterate(
		func(dcl si.ITypeSpec) error {
			if doBefore(dcl) {
				return self.internalGenerate(writer, dcl)
			}
			return nil
		})
	exportName := self.exportDclName(dcl)
	if dcl.Forward() {
		self.writeLineNumber(writer, dcl)
		self.writeComment(writer, "Forward Interface Decl: %v", exportName)
		return nil
	}

	self.writeLineNumber(writer, dcl)
	self.writeComment(writer, "Interface Decl: %v", exportName)
	self.writeLine(writer, true, "type %v interface {", exportName)
	self.incIndent(writer)
	err = multierr.Append(
		err,
		dcl.Iterate(
			func(dcl si.ITypeSpec) error {
				if !doBefore(dcl) {
					return self.internalGenerate(writer, dcl)
				}
				return nil
			}))
	self.decIndent(writer)
	self.writeLine(writer, true, "}")
	return err
}

func (self GolangCodeGenerator) generateOperationDcl(writer io.Writer, dcl si.IOperationDeclarations) error {
	self.writeLineNumber(writer, dcl)
	self.write(writer, true, "%v(", self.exportMemberName(dcl.GetName()))
	if dcl.GetParams() != nil {
		for param := dcl.GetParams(); param != nil; param = param.GetNextParameterDeclarations() {
			self.write(writer, false, "%v %v", self.incomingParamName(param.GetParamName()), self.exportDclName(param.GetParamDeclarationType()))
			if param.GetNextParameterDeclarations() != nil {
				self.write(writer, false, ",")
			}
		}
	}
	self.writeLine(writer, false, ") (%v, error)", self.exportDclName(dcl.GetOperationDeclaratorType()))
	return nil
}

func (self GolangCodeGenerator) generateTypeDcl(writer io.Writer, dcl si.ITypeDeclarator) error {
	self.writeLineNumber(writer, dcl)
	return nil
}

func (self GolangCodeGenerator) generateConstantDcl(writer io.Writer, dcl si.IIdlConstDcl) error {
	self.writeLineNumber(writer, dcl)
	return nil
}

func (self GolangCodeGenerator) generateStructDcl(writer io.Writer, dcl si.IStructType) error {
	self.writeLineNumber(writer, dcl)
	self.writeComment(writer, "Struct Decl: %v", self.exportDclName(dcl))
	return self.generateBaseStructDcl(writer, dcl)
}

func (self GolangCodeGenerator) generateExceptionDcl(writer io.Writer, dcl si.IIdlException) error {
	self.writeLineNumber(writer, dcl)
	self.writeComment(writer, "Exception Decl: %v", self.exportDclName(dcl))
	return self.generateBaseStructDcl(writer, dcl)
}

func (self GolangCodeGenerator) generateAttributeDcl(writer io.Writer, dcl si.IAttributeDcl) error {
	self.writeLineNumber(writer, dcl)
	return nil
}

func (self GolangCodeGenerator) generateBaseStructDcl(writer io.Writer, dcl si.IBaseStructType) error {
	var err error
	exportName := self.exportDclName(dcl)
	var memberList []si.IStructMemberInformation
	if dcl.Members() != nil {
		memberList = dcl.Members().GetMembers()
	}
	self.writeLine(writer, true, "type %v struct {", exportName)
	if dcl.Members() != nil {
		self.incIndent(writer)
		for _, memberInfo := range memberList {
			self.writeLine(writer, true, "%v %v", self.exportMemberName(memberInfo.GetId()), self.exportDclName(memberInfo.GetTypeSpec()))
		}
		self.decIndent(writer)
	}
	self.writeLine(writer, true, "}")
	self.writeLine(writer, true, "")
	self.writeComment(writer, "Constructors")
	self.writeLine(writer, true, "func New%vDefaultPointer() (*%v, error) {", exportName, exportName)
	if dcl.Members() != nil {
		self.incIndent(writer)
		self.write(writer, true, "return &")
		self.writeDefaultValue(writer, dcl, 0)
		self.writeLine(writer, false, ", nil")
		self.decIndent(writer)

	}
	self.writeLine(writer, true, "}")


	self.blankLine(writer)
	self.writeLine(writer, true, "func New%vDefaultValue() (%v, error) {", exportName, exportName)
	if dcl.Members() != nil {
		self.incIndent(writer)
		self.write(writer, true, "return ")
		self.writeDefaultValue(writer, dcl, 0)
		self.writeLine(writer, false, ", nil")
		self.decIndent(writer)
	}
	self.writeLine(writer, true, "}")


	self.blankLine(writer)





	if dcl.Members() != nil {
		self.writeLine(writer, true, "func New%vValue(", exportName)
		self.incIndent(writer)
		for i, memberInfo := range memberList {
			self.write(writer, true, "%v %v", self.incomingParamName(memberInfo.GetId()), self.exportDclName(memberInfo.GetTypeSpec()))
			if i == len(memberList)-1 {

			} else {
				self.writeLine(writer, false, ",")
			}
		}
		self.writeLine(writer, false, ") (%v, error) {", exportName)
		self.writeLine(writer, true, "return %v {", exportName)
		self.incIndent(writer)
		for _, memberInfo := range memberList {
			self.writeLine(writer, true, "%v: %v,", self.exportMemberName(memberInfo.GetId()), self.incomingParamName(memberInfo.GetId()))
		}
		self.decIndent(writer)
		self.writeLine(writer, true, "}, nil")
		self.decIndent(writer)
		self.writeLine(writer, true, "}")
	}
	return err
}

func (self GolangCodeGenerator) writeDefaultValue(writer io.Writer, dcl si.IDeclaredType, count int) {
	switch dcl.GetKind() {
	case si.StructIdlType, si.ExceptionIdlType:
		if baseStruct, ok := dcl.(si.IBaseStructType); ok {
			exportName := self.exportDclName(dcl)
			self.writeLine(writer, false, "%v {", exportName)
			if baseStruct.Members() != nil {
				memberList := baseStruct.Members().GetMembers()
				self.incIndent(writer)
				for _, memberInfo := range memberList {
					self.write(writer, true, "%v: ", self.exportMemberName(memberInfo.GetId()))
					self.writeDefaultValue(writer, memberInfo.GetTypeSpec(), count+1)
				}
				self.decIndent(writer)
			}
			self.write(writer, true, "}")

		}
	case si.RWfloatIdlType:
		self.write(writer, false, "0.0")
	case si.StringIdlType:
		self.write(writer, false, "\"\"")
	case si.WideStringIdlType:
		self.write(writer, false, "\"\"")
	case si.RWVoidType:
		self.write(writer, false, "error")
	case si.RWdoubleIdlType:
		self.write(writer, false, "0.0")
	case si.RWlongRWdoubleIdlType:
		self.write(writer, false, "0")
	case si.RWshortIdlType:
		self.write(writer, false, "0")
	case si.RWlongIdlType:
		self.write(writer, false, "0")
	case si.RWint8IdlType:
		self.write(writer, false, "0")
	case si.RWlongRWlongIdlType:
		self.write(writer, false, "0")
	case si.RWuint8IdlType:
		self.write(writer, false, "0")
	case si.RWunsignedRWshortIdlType:
		self.write(writer, false, "0")
	case si.RWunsignedRWlongIdlType:
		self.write(writer, false, "0")
	case si.RWunsignedRWlongRWlongIdlType:
		self.write(writer, false, "0")
	case si.RWcharIdlType:
		self.write(writer, false, "'0'")
	case si.RWwcharIdlType:
		self.write(writer, false, "'0'")
	case si.RWbooleanIdlType:
		self.write(writer, false, "false")
	case si.RWoctetIdlType:
		self.write(writer, false, "0")
	case si.RWanyIdlType:
		self.write(writer, false, "nil")
	case si.RWObjectIdlType:
		self.write(writer, false, "nil")
	case si.RWint16IdlType:
		self.write(writer, false, "0")
	case si.RWint32IdlType:
		self.write(writer, false, "0")
	case si.RWint64IdlType:
		self.write(writer, false, "0")
	case si.RWuint16IdlType:
		self.write(writer, false, "0")
	case si.RWuint32IdlType:
		self.write(writer, false, "0")
	case si.RWuint64IdlType:
		self.write(writer, false, "0")
	case si.RWEnumIdlType:
		self.write(writer, false, "0")
	case si.RWvaluebaseIdlType:
		self.write(writer, false, "nil")
	default:
		self.writeLine(writer, false, "need default value for '%v'", dcl.GetKind().String())
	}
	if count == 0 {
		self.write(writer, false, "")
	} else {
		self.writeLine(writer, false, ",")
	}
}

func (self GolangCodeGenerator) writeLine(writer io.Writer, useTabs bool, f string, p ...interface{}) {
	tabs := func(tabs bool) string {
		if tabs {
			return strings.Repeat("\t", self.indent)
		}
		return ""
	}(useTabs)
	s := fmt.Sprintf(f, p...)
	data := fmt.Sprintf("%v%v\n", tabs, s)
	if writer != nil {
		_, _ = writer.Write([]byte(data))
	} else {
		fmt.Printf(data)
	}
}

func (self GolangCodeGenerator) write(writer io.Writer, useTabs bool, f string, p ...interface{}) {
	tabs := func(tabs bool) string {
		if tabs {
			return strings.Repeat("\t", self.indent)
		}
		return ""
	}(useTabs)
	s := fmt.Sprintf(f, p...)
	data := fmt.Sprintf("%v%v", tabs, s)
	if writer != nil {
		_, _ = writer.Write([]byte(data))
	} else {
		fmt.Printf(data)
	}
}

func (self GolangCodeGenerator) writeComment(writer io.Writer, f string, p ...interface{}) {
	tabs := strings.Repeat("\t", self.indent)
	s := fmt.Sprintf(f, p...)
	data := fmt.Sprintf("%v//%v\n", tabs, s)
	if writer != nil {
		_, _ = writer.Write([]byte(data))
	} else {
		fmt.Printf(data)
	}

}

func (self GolangCodeGenerator) blankLine(writer io.Writer) {

	data := fmt.Sprintf("\n")
	if writer != nil {
		_, _ = writer.Write([]byte(data))
	} else {
		fmt.Printf(data)
	}

}

func (self *GolangCodeGenerator) incIndent(writer io.Writer) {
	self.indent++
}

func (self *GolangCodeGenerator) decIndent(writer io.Writer) {
	self.indent--
}