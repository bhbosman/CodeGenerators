package yacc

import (
	"bufio"
	"fmt"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scopedObjects"
	"github.com/bhbosman/CodeGenerators/idlgenerator/scoping"
	"github.com/bhbosman/gomock/gomock"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

func TestDefineDef(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
	t.Run("#define Tokenizer", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		mock := NewMockIDefinitionContext(controller)
		mock.EXPECT().ParseDefinition(gomock.Any(), gomock.Any()).MinTimes(1).DoAndReturn(func(s string, b bool) {
			assert.Equal(t, "Tokenizer", s)
			assert.Equal(t, true, b)
		})

		stream := fmt.Sprintf(`
			#define Tokenizer 
			;
		`)
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.CreateWithDefinitionContext("(string test)", reader, mock)
		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 0, result, mockCompleteIdlLexer.LastError())
	})

}
