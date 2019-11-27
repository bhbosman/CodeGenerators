package yacc

//go:generate goyacc -o completeIdl.go  -p "CompleteIdl"  completeIdl.y

import (
	"bufio"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scoping"
	"github.com/bhbosman/gomock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestIfEnumsAreParsing(t *testing.T) {
	//defaultTypeService := scoping.NewDefaultTypeService()
	//scopingContextFactory := scoping.NewScopingContextFactory(defaultTypeService)
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext("", scoping.NewDefaultTypeService(), nil))
	type streamTest struct {
		name   string
		stream string
		pass   bool
	}
	data := []streamTest{
		{
			name:   "test enum 01",
			stream: "enum enumName {a,b,c};",
			pass:   true,
		},
		{
			name:   "test enum 02",
			stream: "enum EnumName{a};",
			pass:   true,
		},
		{
			name:   "test enum 03",
			stream: "enum enumName {a,b};",
			pass:   true,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			reader := bufio.NewReader(strings.NewReader(d.stream))
			idlExprLex, _ := factory.Create("(string test)", reader)
			result := CompleteIdlParse(idlExprLex)
			assert.Equal(t, 0, result)
		})
	}
}
