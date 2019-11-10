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

func TestForwardInterfaceDecl(t *testing.T) {
	factory := NewCompleteIdlLexerFactoryImpl(
		log.New(os.Stdout, "test", 0),
		scopedObjects.NewNextNumber(),
		scoping.NewScopingContext(scoping.NewDefaultTypeService(), nil))
	t.Run("Forward Interface Decl", func(t *testing.T) {
		stream := fmt.Sprintf(`
			interface IABC;
		`)
		controller := gomock.NewController(t)
		defer controller.Finish()
		reader := bufio.NewReader(strings.NewReader(stream))
		actual, _ := factory.Create("(string test)", reader)

		mockCompleteIdlLexer := NewMockCompleteIdlLexerInstanceWrapper(controller, actual)
		mockCompleteIdlLexer.EXPECT().CreateInterfaceDcl(gomock.Any(), "IABC", gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
		result := CompleteIdlParse(mockCompleteIdlLexer)
		assert.Equal(t, 0, result, mockCompleteIdlLexer.LastError())
	})
}
