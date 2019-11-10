package publishGo

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/Common"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/Extensions"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/yacc"
	"io"
)

type publishEnum struct {
	data *yacc.EnumDecl
}

func (self publishEnum) ExportDefinition(writer io.StringWriter) {
	_, _ = writer.WriteString(fmt.Sprintf("type %s byte \n", self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("//noinspection ALL\n"))
	_, _ = writer.WriteString(fmt.Sprintf("const (\n"))
	for index, decl := range self.data.Decls {
		defaultValue := decl.DefaultValue()

		if defaultValue == nil {
			if index == 0 {
				_, _ = writer.WriteString(fmt.Sprintf("\t%v_%v %v = iota\n", self.data.Identifier, decl.Identifier(), self.data.Identifier))
			} else {
				_, _ = writer.WriteString(fmt.Sprintf("\t%v_%v\n", self.data.Identifier, decl.Identifier()))
			}
		} else {
			value := defaultValue.Value()
			switch v := value.(type) {
			case *byte:
			case byte:
				_, _ = writer.WriteString(fmt.Sprintf(
					"\t%v_%v  = %v // char value: '%v'\n ",
					self.data.Identifier,
					decl.Identifier(),
					v,
					string(v)))

			default:
				_, _ = writer.WriteString(fmt.Sprintf(
					"\t%v_%v  = %v // default value: byte(%v)\n ",
					self.data.Identifier,
					decl.Identifier(),
					v,
					v))

			}

		}
	}
	_, _ = writer.WriteString(fmt.Sprintf(")\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func (self publishEnum) Export(
	writer io.StringWriter) {
	self.ExportDefinition(writer)
	self.ExportFactoryDefinition(writer)

	typeNamePrefix := self.data.Identifier
	typeCode := CalculateCrc(typeNamePrefix)
	_, _ = writer.WriteString(fmt.Sprintf("// %v Declaration TypeCode: 0x%08x\n", typeNamePrefix, typeCode))

	_ = self.GenerateWriteFunction(writer, typeNamePrefix, typeCode)

}

func (self publishEnum) GenerateWriteFunction(
	writer io.StringWriter,
	typeNamePrefix string, typeCode uint32) error {

	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		returnType := Extensions.TypeValueHelper.TypeValueForDefinedType(self.data)
		_, _ = writer.WriteString(fmt.Sprintf("// %v writer \n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("func Write_%v(stream Streams.I%vWriter, value %v) (int, error) {\n", typeNamePrefix, "Mitch", returnType))
		_, _ = writer.WriteString(fmt.Sprintf("\treturn stream.Write_byte(byte(value))\n"))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func (self publishEnum) ExportFactoryDefinition(writer io.StringWriter) {
	_, _ = writer.WriteString(fmt.Sprintf("type %sFactoryType struct {\n", self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
	_, _ = writer.WriteString(fmt.Sprintf("var %sFactory %sFactoryType = %sFactoryType{}\n", self.data.Identifier, self.data.Identifier, self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("func (self %sFactoryType) ReadValue(stream Streams.IMitchReader) (%v, int, error) {\n", self.data.Identifier, self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("\tvar n int\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tvar b byte\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tvar err error\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tb, n, err = stream.Read_byte()\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t\treturn %v(0), 0, err\n", self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\treturn %v(b), n, nil\n", self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))

}

func NewPublishEnum(data *yacc.EnumDecl) *publishEnum {
	return &publishEnum{
		data: data,
	}
}
