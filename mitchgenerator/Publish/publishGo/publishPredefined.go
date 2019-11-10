package publishGo

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/Extensions"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
	"hash/crc32"
	"io"
)

func CalculateCrc(identifier string) uint32 {
	crc32q := crc32.MakeTable(0xD5828281)
	return crc32.Checksum([]byte(identifier), crc32q)
}

type publishPredefined struct {
	data interfaces.IDefinitionDeclaration
}

func (self *publishPredefined) Export(
	writer io.StringWriter) {

	if self.data.Predefined() {
		if self.data.Kind() == interfaces.Invalid {
			return
		}
		typeNamePrefix := self.data.Kind().String()
		typeCode := CalculateCrc(typeNamePrefix)

		_, _ = writer.WriteString(fmt.Sprintf("// %v Declaration TypeCode: 0x%08x\n", typeNamePrefix, typeCode))

		self.GenerateWriteFunction(
			writer,
			GenerateWriteFunctionParams{

				typeNamePrefix: typeNamePrefix,
				typeCode:       typeCode})
		self.GenerateReadFunction(
			writer,
			GenerateReadFunctionParams{
				typeNamePrefix: typeNamePrefix,
				typeCode:       typeCode,
				defaultValue:   self.data.DefaultValue()})
	}
}

type GenerateReadFunctionParams struct {
	typeNamePrefix string
	typeCode       uint32
	defaultValue   interface{}
}

func (self *publishPredefined) GenerateReadFunction(
	writer io.StringWriter,

	params GenerateReadFunctionParams) {

	returnType := Extensions.TypeValueHelper.TypeValueForDefinedType(self.data)
	_, _ = writer.WriteString(fmt.Sprintf("// %v reader\n", params.typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("func Read_%v(stream Streams.I%vReader) (%v, int, error) {\n", params.typeNamePrefix, "Mitch", returnType))
	_, _ = writer.WriteString(fmt.Sprintf("\treturn stream.Read_%v()\n", params.typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

type GenerateWriteFunctionParams struct {
	typeNamePrefix string
	typeCode       uint32
}

func (self *publishPredefined) GenerateWriteFunction(writer io.StringWriter, params GenerateWriteFunctionParams) {
	returnType := Extensions.TypeValueHelper.TypeValueForDefinedType(self.data)
	_, _ = writer.WriteString(fmt.Sprintf("// %v writer \n", params.typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("func Write_%v(stream Streams.I%vWriter, value %v) (int, error) {\n", params.typeNamePrefix, "Mitch", returnType))
	_, _ = writer.WriteString(fmt.Sprintf("\treturn stream.Write_%v(value)\n", params.typeNamePrefix))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\n"))
}

func NewPublishPredefined(data interfaces.IDefinitionDeclaration) *publishPredefined {
	return &publishPredefined{
		data: data,
	}
}
