package publishGo

import (
	"fmt"
	"github.com/bhbosman/CodeGenerators/Common"

	"github.com/bhbosman/CodeGenerators/mitchgenerator/MitchDefinedTypes"
	//"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
	"io"
)

type publishBitField struct {
	data       *MitchDefinedTypes.MitchBitField
	Identifier string
}

func (self publishBitField) Export(writer io.StringWriter) {
	_ = self.ExportDefinition(writer)
	_ = self.ExportDefaultConstructor(writer)
	typeNamePrefix := self.Identifier
	typeCode := CalculateCrc(typeNamePrefix)
	_, _ = writer.WriteString(fmt.Sprintf("// %v Declaration TypeCode: 0x%08x\n", typeNamePrefix, typeCode))

	_ = self.GenerateWriteFunction(writer, typeNamePrefix, typeCode)
	_ = self.GenerateReadFunction(writer, typeNamePrefix, typeCode)
}

func (self *publishBitField) ExportDefinition(writer io.StringWriter) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		_, _ = writer.WriteString(fmt.Sprintf("type %s struct{\n", self.Identifier))
		_, _ = writer.WriteString(fmt.Sprintf("\tFlags byte\n"))

		//if self.data. BitsUsed&0x01 == 0x01 {
		//	_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 00\n", self.data.BitField00))
		//}
		//if self.data.BitsUsed&0x02 == 0x02 {
		//	_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 01\n", self.data.BitField01))
		//}
		//if self.data.BitsUsed&0x04 == 0x04 {
		//	_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 02\n", self.data.BitField02))
		//}
		//if self.data.BitsUsed&0x08 == 0x08 {
		//	_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 03\n", self.data.BitField03))
		//}
		//if self.data.BitsUsed&0x10 == 0x10 {
		//	_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 04\n", self.data.BitField04))
		//}
		//if self.data.BitsUsed&0x20 == 0x20 {
		//	_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 05\n", self.data.BitField05))
		//}
		//if self.data.BitsUsed&0x40 == 0x40 {
		//	_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 06\n", self.data.BitField06))
		//}
		//if self.data.BitsUsed&0x80 == 0x80 {
		//	_, _ = writer.WriteString(fmt.Sprintf("\t%v bool // Bit 07\n", self.data.BitField07))
		//}
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func (self *publishBitField) ExportDefaultConstructor(writer io.StringWriter) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		_, _ = writer.WriteString(fmt.Sprintf("func New%v()%v {\n", self.Identifier, self.Identifier))
		_, _ = writer.WriteString(fmt.Sprintf("\treturn %v{}\n", self.Identifier))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func (self publishBitField) GenerateReadFunction(writer io.StringWriter, typeNamePrefix string, typeCode uint32) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		_, _ = writer.WriteString(fmt.Sprintf("// %v reader\n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("func Read_%v(stream Streams.I%vReader) (value %v, byteCount int, err error) {\n", typeNamePrefix, "Mitch", self.Identifier))
		_, _ = writer.WriteString(fmt.Sprintf("\tvalue = New%v()\n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("\tvalue.Flags, byteCount, err = stream.Read_byte()\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\treturn value, byteCount, err\n"))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func (self publishBitField) GenerateWriteFunction(writer io.StringWriter, typeNamePrefix string, typeCode uint32) error {
	return Common.ErrorListFactory.NewErrorListFunc(func(errorList Common.IErrorList) {
		_, _ = writer.WriteString(fmt.Sprintf("// %v writer \n", typeNamePrefix))
		_, _ = writer.WriteString(fmt.Sprintf("func Write_%v(stream Streams.I%vWriter, value %v) (byteCount int, err error) {\n", typeNamePrefix, "Mitch", self.Identifier))
		_, _ = writer.WriteString(fmt.Sprintf("\treturn stream.Write_byte(byte(value.Flags))\n"))
		_, _ = writer.WriteString(fmt.Sprintf("}\n"))
		_, _ = writer.WriteString(fmt.Sprintf("\n"))
	})
}

func NewPublishBitField(data *MitchDefinedTypes.MitchBitField, identifier string) *publishBitField {
	return &publishBitField{
		data:       data,
		Identifier: identifier,
	}
}
