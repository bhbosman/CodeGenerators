package publishGo

import (
	"github.com/bhbosman/CodeGenerators/mitchgenerator/MitchDefinedTypes"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/yacc"
	"io"
)

type publishTypeDecl struct {
	data *yacc.TypeDeclarator
}

func (self *publishTypeDecl) ExportDefinition(writer io.StringWriter) {
	//_, _ = writer.WriteString(fmt.Sprintf("type %s struct {\n", self.data.Identifier))
	//for _, member := range self.data.Members {
	//	defaultValue := member.Declarator.DefaultValue()
	//
	//	if defaultValue == nil {
	//		_, _ = writer.WriteString(fmt.Sprintf(
	//			"\t%v %v\n",
	//			member.Declarator.Identifier(),
	//			member.DefinedType.GetName()))
	//	} else {
	//
	//		ss := defaultValue.ValueKind().String()
	//		_, _ = writer.WriteString(fmt.Sprintf(
	//			"\t%v %v // default value: %v(%v)\n ",
	//			member.Declarator.Identifier(),
	//			member.DefinedType.GetName(),
	//			ss,
	//			GetExportValue(defaultValue)))
	//	}
	//}
	//_, _ = writer.WriteString(fmt.Sprintf("}\n"))
	//_, _ = writer.WriteString(fmt.Sprintf("\n"))}

}

func (self *publishTypeDecl) Export(writer io.StringWriter) {
	if definition, ok := self.data.DefinedTyped.(*MitchDefinedTypes.MitchBitField); ok {
		publish := NewPublishBitField(definition, self.data.Declarator.Identifier())
		publish.Export(writer)
	}
}

func NewpublishTypeDecl(data *yacc.TypeDeclarator) *publishTypeDecl {
	return &publishTypeDecl{
		data: data,
	}
}
