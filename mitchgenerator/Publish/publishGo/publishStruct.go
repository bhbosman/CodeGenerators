package publishGo

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/Common"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/Extensions"
	. "github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/yacc"
	"io"
)

type publishStruct struct {
	data *yacc.MitchMessageDefinition
}

type LineWriter struct {
	errorList Common.IErrorList
	writer    io.StringWriter
}

func (receiver LineWriter) Write(format string, a ...interface{}) {
	_, e := receiver.writer.WriteString(fmt.Sprintf(format, a...))
	receiver.errorList.Add(e)
}

func NewLineWriter(writer io.StringWriter, errorList Common.IErrorList) LineWriter {
	return LineWriter{
		errorList: errorList,
		writer:    writer,
	}
}

func (self *publishStruct) Export(writer io.StringWriter) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		lineWriter := NewLineWriter(writer, errorList)
		typeNamePrefix := self.data.Identifier
		typeCode := CalculateCrc(typeNamePrefix)
		lineWriter.Write("// %v Declaration TypeCode: 0x%08x\n", typeNamePrefix, typeCode)
		errorList.Add(self.ExportDefinition(writer, typeNamePrefix))
		errorList.Add(self.ExportFactoryDefinition(writer, typeNamePrefix))
		errorList.Add(self.GenerateWriteFunction(writer, typeNamePrefix, typeCode))
	})
}

func (self *publishStruct) ExportFactoryDefinition(writer io.StringWriter, typeNamePrefix string) error {
	_, _ = writer.WriteString(fmt.Sprintf("type %sFactoryType struct {\n", self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))

	_, _ = writer.WriteString(fmt.Sprintf("var %sFactory %sFactoryType = %sFactoryType{}\n", self.data.Identifier, self.data.Identifier, self.data.Identifier))
	if self.data.HasMessageInformation() {
		_, _ = writer.WriteString(fmt.Sprintf("const %s_MessageType int = 0x%x//%v\n", self.data.Identifier, self.data.MessageType, self.data.MessageType))
		_, _ = writer.WriteString(fmt.Sprintf("const %s_MessageLength uint16 = %v\n", self.data.Identifier, self.data.MessageLength))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	}

	_, _ = writer.WriteString(fmt.Sprintf("func(self %sFactoryType) New() (*%s, error) {\n", self.data.Identifier, self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("\treturn &%s{}, nil\n", self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))

	_, _ = writer.WriteString(fmt.Sprintf("func (self %sFactoryType) ReadMessageHeader(instance *%s, stream Streams.IMitchReader) (byteCount int, err error) {\n", self.data.Identifier, self.data.Identifier))
	if self.data.HasMessageInformation() {
		_, _ = writer.WriteString(fmt.Sprintf("\tvalue, n, err := stream.Read_uint16()\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\tif value != 0 && value != 0x%x {\n", self.data.MessageLength))
		_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, fmt.Errorf(\"message length incorrect. For Message %v was expected 0x%x, but 0x%%x was found.)\", value)\n", self.data.Identifier, self.data.MessageLength))
		_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\tb, n, err := stream.Read_byte()\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\tif b != 0x%x {\n", self.data.MessageType))
		_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, fmt.Errorf(\"message type numbers does not match up. For Message %v was expected 0x%x, but 0x%%x was found.)\", b)\n", self.data.Identifier, self.data.MessageType))
		_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\treturn byteCount, nil\n"))

	} else {
		_, _ = writer.WriteString(fmt.Sprintf("\treturn 0, nil\n"))
	}
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))

	_, _ = writer.WriteString(fmt.Sprintf("func (self %sFactoryType) ReadMessageData(instance *%s, stream Streams.IMitchReader) (byteCount int, err error) {\n", self.data.Identifier, self.data.Identifier))
	_, _ = writer.WriteString(fmt.Sprintf("\tvar n int\n"))
	for index, item := range self.data.Members {
		if item.DefinedType == nil {
			return fmt.Errorf("ddd")
		}

		_, _ = writer.WriteString(fmt.Sprintf("\t// Index: %v, StructMember GetName: %v, Type: %v \n", index, item.Declarator.Identifier(), item.DefinedType.GetName()))
		switch item.DefinedType.Kind() {
		case MitchAlpha:
			_, seqCount := item.DefinedType.GetSequenceCount()
			_, _ = writer.WriteString(fmt.Sprintf("\tinstance.%v, n, err = stream.Read_%v(%v)\n", item.Declarator.Identifier(), item.DefinedType.GetStreamFunctionName(), seqCount))
			_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\n"))
			break
		case Enum:
			//_, _ = writer.WriteString(fmt.Sprintf("\ttemp_%v, n, err := stream.Read_byte()\n", item.Declarator.Identifier()))
			_, _ = writer.WriteString(fmt.Sprintf("\tinstance.%v, n, err = %vFactory.ReadValue(stream)\n", item.Declarator.Identifier(), item.DefinedType.GetName()))
			_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\n"))
			break
		case MitchBitField:
			if typeDecl, ok := item.DefinedType.(ITypeDeclaration); ok {
				_, _ = writer.WriteString(fmt.Sprintf("\tinstance.%v, n, err = Read_%v(stream)\n", item.Declarator.Identifier(), typeDecl.GetDeclarator().Identifier()))
				_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
				_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
				_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
				_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
				_, _ = writer.WriteString(fmt.Sprintf("\n"))
			} else {
				panic("DDDD")
			}
		default:
			_, _ = writer.WriteString(fmt.Sprintf("\tinstance.%v, n, err = stream.Read_%v()\n", item.Declarator.Identifier(), item.DefinedType.GetStreamFunctionName()))
			_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\n"))

		}
	}
	_, _ = writer.WriteString(fmt.Sprintf("\treturn byteCount, nil\n"))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))

	_, _ = writer.WriteString(fmt.Sprintf("func (self %sFactoryType) ReadMessageInFull(instance *%s, stream Streams.IMitchReader) (byteCount int, err error) {\n", self.data.Identifier, self.data.Identifier))

	_, _ = writer.WriteString(fmt.Sprintf("\tn, err := self.ReadMessageHeader(instance, stream)\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tn, err = self.ReadMessageData(instance, stream)\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\treturn byteCount, nil\n"))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))

	return nil

}

func (self *publishStruct) ExportDefinition(writer io.StringWriter, typeNamePrefix string) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		_, _ = writer.WriteString(fmt.Sprintf("type %s struct {\n", self.data.Identifier))
		for _, member := range self.data.Members {
			defaultValue := member.Declarator.DefaultValue()
			if member.DefinedType == nil {
				errorList.Add(fmt.Errorf("type is null"))
				return
			}
			returnType := Extensions.TypeValueHelper.TypeValueForDefinedType(member.DefinedType)
			if defaultValue == nil {
				_, _ = writer.WriteString(fmt.Sprintf(
					"\t%v %v   `json:\"%v\" xml:\"%v,attr\"`  \n",
					member.Declarator.Identifier(),
					returnType,
					member.Declarator.Identifier(),
					member.Declarator.Identifier()))
			} else {
				ss := defaultValue.ValueKind().String()
				_, _ = writer.WriteString(fmt.Sprintf(
					"\t%v %v // default value: %v(%v)\n ",
					member.Declarator.Identifier(),
					returnType,
					ss,
					defaultValue.Value()))
			}
		}
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func (self publishStruct) GenerateWriteFunction(writer io.StringWriter, typeNamePrefix string, typeCode uint32) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		returnType := "*" + typeNamePrefix
		_, _ = writer.WriteString(fmt.Sprintf("// %v writer \n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("func Write_%v(stream Streams.I%vWriter, value %v) (byteCount int, err error) {\n", typeNamePrefix, "Mitch", returnType))
		_, _ = writer.WriteString(fmt.Sprintf("\tvar n int \n"))

		if self.data.HasMessageInformation() {
			_, _ = writer.WriteString(fmt.Sprintf("\tn, err = stream.Write_uint16(%v)\n", self.data.MessageLength))
			_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\tn, err = stream.Write_byte(%v)\n", self.data.MessageType))
			_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
		}
		for index, item := range self.data.Members {
			if item.DefinedType == nil {
				errorList.Add(fmt.Errorf("defined type is null"))
				return
			}
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t// Index: %v, StructMember GetName: %v, Type: %v \n", index, item.Declarator.Identifier(), item.DefinedType.GetName()))
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t//\n"))
			switch item.DefinedType.Kind() {
			case MitchAlpha:
				_, seqCount := item.DefinedType.GetSequenceCount()
				_, _ = writer.WriteString(fmt.Sprintf("\tn, err = stream.Write_%v(value.%v, %v)\n", item.DefinedType.GetStreamFunctionName(), item.Declarator.Identifier(), seqCount))

			case Enum:
				_, _ = writer.WriteString(fmt.Sprintf("\tn, err = stream.Write_byte(byte(value.%v))\n", item.Declarator.Identifier()))
				break
			case MitchBitField:
				if typeDecl, ok := item.DefinedType.(ITypeDeclaration); ok {
					_, _ = writer.WriteString(fmt.Sprintf("\tn, err = Write_%v(stream, value.%v)\n", typeDecl.GetDeclarator().Identifier(), item.Declarator.Identifier()))
				} else {
					panic("DDDD")
				}
			default:
				_, _ = writer.WriteString(fmt.Sprintf("\tn, err = stream.Write_%v(value.%v)\n", item.DefinedType.GetStreamFunctionName(), item.Declarator.Identifier()))
			}
			_, _ = writer.WriteString(fmt.Sprintf("\tif err != nil {\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t\treturn 0, err\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
			_, _ = writer.WriteString(fmt.Sprintf("\tbyteCount += n\n"))
		}
		_, _ = writer.WriteString(fmt.Sprintf("\treturn byteCount, nil\n"))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func NewPublishStruct(data *yacc.MitchMessageDefinition) *publishStruct {
	return &publishStruct{
		data: data,
	}
}
