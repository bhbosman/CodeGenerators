package publishGo

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/interfaces"
	"github.com/bhbosman/CodeGenerators/mitchgenerator/yacc"
)

type MessageGenerator struct {
	DeclaredTypes   []interfaces.IDefinitionDeclaration
	typeInformation interfaces.IBaseTypeInformation
}

func (self *MessageGenerator) Export(writer *bufio.Writer, information interfaces.IBaseTypeInformation) error {
	structDefinitionCount := 0
	for _, declaredType := range self.DeclaredTypes {
		if _, ok := declaredType.(*yacc.MitchMessageDefinition); ok {
			structDefinitionCount++
		}
	}
	_, _ = writer.WriteString(fmt.Sprintf("type GenerateMessages struct {\n"))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))

	_, _ = writer.WriteString(fmt.Sprintf("func (self *GenerateMessages) PublishMessage() (interface{}, error) {\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\tindex := rand.Int()/ %d\n", structDefinitionCount))

	_, _ = writer.WriteString(fmt.Sprintf("\tswitch index {\n"))

	structDefinitionIndex := 0
	for _, declaredType := range self.DeclaredTypes {
		if structDefinition, ok := declaredType.(*yacc.MitchMessageDefinition); ok {
			_, _ = writer.WriteString(fmt.Sprintf("\tcase %d:\n", structDefinitionIndex))
			_, _ = writer.WriteString(fmt.Sprintf("\t\treturn New%v(), nil\n", structDefinition.Identifier))
			structDefinitionIndex++
		}
	}
	_, _ = writer.WriteString(fmt.Sprintf("\t}\n"))
	_, _ = writer.WriteString(fmt.Sprintf("\t return nil, fmt.Errorf(\"no message selected\\n\")\n"))
	_, _ = writer.WriteString(fmt.Sprintf("}\n"))

	return nil
}

func NewMessageGenerator(DeclaredTypes []interfaces.IDefinitionDeclaration, typeInformation interfaces.IBaseTypeInformation) *MessageGenerator {
	return &MessageGenerator{
		DeclaredTypes:   DeclaredTypes,
		typeInformation: typeInformation,
	}
}
